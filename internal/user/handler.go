package user

import (
	"avitoIntershipBackend/internal/handlers"
	"avitoIntershipBackend/internal/masterBalance"
	"avitoIntershipBackend/pkg/logging"
	"github.com/gin-gonic/gin"
	"strings"
)

var _ handlers.Handler = &handler{}

const (
	billingURL          = "/billing"
	freezeURL           = "/moneyFreeze"
	acceptURL           = "/moneyAccept"
	rejectURL           = "/moneyReject"
	userURL             = "/users/:id"
	userTransactionsURL = "/transactions/:userid/:pageNum/:sortSum/:sortDate"
	reportURL           = "/report/:month/:year"
)

type handler struct {
	logger      logging.Logger
	userService BusinessLogic
}

func NewHandler(logger logging.Logger, serviceUser BusinessLogic) handlers.Handler {
	return &handler{
		logger:      logger,
		userService: serviceUser,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST(billingURL, h.AddBilling)
	router.POST(freezeURL, h.FreezeMoney)
	router.POST(acceptURL, h.AcceptMoney)
	router.POST(rejectURL, h.RejectMoney)
	router.GET(userURL, h.GetUserBalance)
	router.GET(reportURL, h.GetReport)
	router.GET(userTransactionsURL, h.GetUserTransactions)

}

func (h *handler) AddBilling(c *gin.Context) {
	var user User
	var err error
	if err = c.BindJSON(&user); err != nil {
		c.AbortWithStatus(500)
	}
	err = h.userService.Billing(c, &user)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect \"balanace\" parametr in request") {
			c.AbortWithStatus(400)
		}
		c.AbortWithStatus(500)
	}
	c.JSON(200, user)
}

func (h *handler) FreezeMoney(c *gin.Context) {
	var masterReq masterBalance.MasterBalance
	var err error
	if err = c.BindJSON(&masterReq); err != nil {
		c.AbortWithStatus(500)
	}
	err = h.userService.ReserveMoney(c, &masterReq)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect \"balanace\" parametr in request") {
			c.AbortWithStatus(400)

		}
		if strings.Contains(err.Error(), "lack of money for service is") {
			c.AbortWithStatus(204)

		}
		c.AbortWithStatus(500)

	}
	c.String(200, "reserved bill id %d", masterReq.ID)
}

func (h *handler) AcceptMoney(c *gin.Context) {
	var masterReq masterBalance.MasterBalance
	var err error
	if err = c.BindJSON(&masterReq); err != nil {
		c.AbortWithStatus(500)
	}
	err = h.userService.AcceptMoney(c, &masterReq)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect \"balanace\" parametr in request") {
			c.AbortWithStatus(400)

		}
		c.AbortWithStatus(500)

	}
	c.Status(200)
}

func (h *handler) RejectMoney(c *gin.Context) {
	var masterReq masterBalance.MasterBalance
	var err error
	if err = c.BindJSON(&masterReq); err != nil {
		c.AbortWithStatus(500)
	}
	err = h.userService.RejectMoney(c, &masterReq)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect \"balanace\" parametr in request") {
			c.AbortWithStatus(400)

		}
		c.AbortWithStatus(500)

	}
	c.Status(200)
}

func (h *handler) GetUserBalance(c *gin.Context) {
	var user User
	var err error

	id := c.Params.ByName("id")

	user, err = h.userService.GetBalance(c, id)
	if err != nil {
		if strings.Contains(err.Error(), "no users with such id") {
			c.AbortWithStatus(404)
		}
		c.AbortWithStatus(500)
	}
	c.JSON(200, user)
}

func (h *handler) GetReport(c *gin.Context) {
	var linkToReport string
	var err error

	month := c.Params.ByName("month")
	year := c.Params.ByName("year")

	linkToReport, err = h.userService.Report(c, month, year)
	if err != nil {
		if strings.Contains(err.Error(), "no transactions for this period") {
			c.AbortWithStatus(404)
		}
		c.AbortWithStatus(500)
	}
	c.String(200, "%s", linkToReport)
}

func (h *handler) GetUserTransactions(c *gin.Context) {
	var transactionsList string
	var err error

	id := c.Params.ByName("userid")
	pageNum := c.Params.ByName("pageNum")
	sortSum := c.Params.ByName("sortSum")
	sortDate := c.Params.ByName("sortDate")

	transactionsList, err = h.userService.GetUserTransactions(c, id, pageNum, sortSum, sortDate)
	if err != nil {
		if strings.Contains(err.Error(), "no transactions for user") {
			c.AbortWithStatus(404)
		}
		if strings.Contains(err.Error(), "incorrect input parametrs") ||
			strings.Contains(err.Error(), "parametrs equal 0") {
			c.AbortWithStatus(400)
		}
		c.AbortWithStatus(500)
	}
	c.Data(200, "text/html; charset=utf-8", []byte(transactionsList))
}
