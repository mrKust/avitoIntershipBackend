package user

import (
	"avitoIntershipBackend/internal/handlers"
	"avitoIntershipBackend/pkg/logging"
	"github.com/gin-gonic/gin"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(usersURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

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
