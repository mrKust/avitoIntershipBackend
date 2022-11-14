package user

import (
	"avitoIntershipBackend/internal/handlers"
	"avitoIntershipBackend/internal/masterBalance"
	"avitoIntershipBackend/pkg/logging"
	"fmt"
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

// AddBilling godoc
// @Summary      Add money to user's balance
// @Description  Add money to user's balance with billing systems (visa/mastercard)
// @Tags         accounts, billings
// @Accept       json
// @Produce      json
// @Param        message  body  User  true  "Users Info"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /billing [post]
func (h *handler) AddBilling(c *gin.Context) {
	var user User
	var err error
	if err = c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	err = h.userService.Billing(c, &user)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect balanace parametr in request") {
			c.JSON(400, gin.H{"err message": err.Error()})
			return
		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	c.JSON(200, user)
}

// FreezeMoney godoc
// @Summary      Reserves money
// @Description  Reserves money from user balance to special master account
// @Tags         accounts, reserve
// @Accept       json
// @Produce      json
// @Param        message  body  masterBalance.CreateDTO  true  "Request to freeze money"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /moneyFreeze [post]
func (h *handler) FreezeMoney(c *gin.Context) {
	var masterReq masterBalance.MasterBalance
	var err error
	if err = c.BindJSON(&masterReq); err != nil {
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	err = h.userService.ReserveMoney(c, &masterReq)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect balanace parametr in request") {
			c.JSON(400, gin.H{"err message": err.Error()})
			return

		}
		if strings.Contains(err.Error(), "lack of money for service is") {
			c.JSON(404, gin.H{"err message": err.Error()})
			return

		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return

	}
	c.JSON(200, gin.H{"res message :": fmt.Sprintf("reserved bill id %d", masterReq.ID)})
}

// AcceptMoney godoc
// @Summary      Accepts money
// @Description  Accept money from master balance when service is done
// @Tags         accounts, reserve
// @Accept       json
// @Produce      json
// @Param        message  body  masterBalance.CreateDTO  true  "Request to freeze money"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /moneyAccept [post]
func (h *handler) AcceptMoney(c *gin.Context) {
	var masterReq masterBalance.MasterBalance
	var err error
	if err = c.BindJSON(&masterReq); err != nil {
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	err = h.userService.AcceptMoney(c, &masterReq)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect balanace parametr in request") {
			c.JSON(400, gin.H{"err message": err.Error()})
			return
		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return

	}
	c.Status(200)
}

// RejectMoney godoc
// @Summary      Rejects money
// @Description  Return money to user when payment for service is rejected
// @Tags         accounts, reserve, reject
// @Accept       json
// @Produce      json
// @Param        message  body  masterBalance.CreateDTO  true  "Request to freeze money"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /moneyReject [post]
func (h *handler) RejectMoney(c *gin.Context) {
	var masterReq masterBalance.MasterBalance
	var err error
	if err = c.BindJSON(&masterReq); err != nil {
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	err = h.userService.RejectMoney(c, &masterReq)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect balanace parametr in request") {
			c.JSON(400, gin.H{"err message": err.Error()})
			return
		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"res message": "transaction canceled"})
}

// GetUserBalance godoc
// @Summary      Returns user balance
// @Description  Return user account with his balance
// @Tags         accounts, balance
// @Accept       json
// @Produce      json
// @Param		 id    path      int     true  "User's id"
// @Success      200
// @Failure      404
// @Failure      500
// @Router       /users/{id} [get]
func (h *handler) GetUserBalance(c *gin.Context) {
	var user User
	var err error

	id := c.Params.ByName("id")

	user, err = h.userService.GetBalance(c, id)
	if err != nil {
		if strings.Contains(err.Error(), "no users with such id") {
			c.JSON(404, gin.H{"err message": err.Error()})
			return
		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	c.JSON(202, user)
}

// GetReport godoc
// @Summary      Returns report for date range
// @Description  Return link to report.csv file with money for every service
// @Tags         accounts, balance, report
// @Accept       json
// @Produce      json
// @Param		 month    path      int     true  "Needed month for report"
// @Param		 year     path      int     true  "Needed year for report"
// @Success      200
// @Failure      404
// @Failure      500
// @Router       /report/{month}/{year} [get]
func (h *handler) GetReport(c *gin.Context) {
	var linkToReport string
	var err error

	month := c.Params.ByName("month")
	year := c.Params.ByName("year")

	linkToReport, err = h.userService.Report(c, month, year)
	if err != nil {
		if strings.Contains(err.Error(), "no transactions for this period") {
			c.JSON(404, gin.H{"err message": err.Error()})
			return
		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"link to report": linkToReport})
}

// GetUserTransactions godoc
// @Summary      Returns info about user transactions
// @Description  Return text with history of transactions
// @Tags         accounts, balance
// @Accept       json
// @Produce      json
// @Param		 userid   path      int     true  "id of needed user"
// @Param		 pageNum  path      int     true  "number of searching page"
// @Param		 sortSum  path      string     true  "Parameter for sort by sum (asc/desc)"
// @Param		 sortDate  path     string     true  "Parameter for sort by date (asc/desc)"
// @Success      200
// @Failure      404
// @Failure      404
// @Failure      500
// @Router       /transactions/{userid}/{pageNum}/{sortSum}/{sortDate} [get]
func (h *handler) GetUserTransactions(c *gin.Context) {
	var transactionsList []string
	var err error

	id := c.Params.ByName("userid")
	pageNum := c.Params.ByName("pageNum")
	sortSum := c.Params.ByName("sortSum")
	sortDate := c.Params.ByName("sortDate")

	transactionsList, err = h.userService.GetUserTransactions(c, id, pageNum, sortSum, sortDate)
	if err != nil {
		if strings.Contains(err.Error(), "no transactions for user") {
			c.JSON(404, gin.H{"err message": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "incorrect input parametrs") ||
			strings.Contains(err.Error(), "parametrs equal 0") {
			c.JSON(400, gin.H{"err message": err.Error()})
			return
		}
		c.JSON(500, gin.H{"err message": err.Error()})
		return
	}
	c.JSON(200, transactionsList)
}
