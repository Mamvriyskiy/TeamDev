package handler

import (
	"github.com/Mamvriyskiy/TeamDev/pkg/service"
	"github.com/gin-gonic/gin"
)

const signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	auth.POST("/register", h.RegisterUser)
	auth.POST("/profile", h.ProfileUser)
	auth.POST("/social", h.AddSocialUser)

	// logger.Log("Info", "", "Create router", nil)

	return router
}
