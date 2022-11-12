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
	billingURL = "/billing"
	freezeURL  = "/moneyFreeze"
	usersURL   = "/users"
	userURL    = "/users/:id"
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
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(usersURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

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

func (h *handler) GetList(c *gin.Context) {
	c.AbortWithStatus(404)
}

func (h *handler) CreateUser(c *gin.Context) {
	c.AbortWithStatus(404)
}
func (h *handler) GetUserByID(c *gin.Context) {
	c.AbortWithStatus(404)
}
func (h *handler) UpdateUser(c *gin.Context) {
	c.AbortWithStatus(404)
}
func (h *handler) PartiallyUpdateUser(c *gin.Context) {
	c.AbortWithStatus(404)
}
func (h *handler) DeleteUser(c *gin.Context) {
	c.AbortWithStatus(404)
}
