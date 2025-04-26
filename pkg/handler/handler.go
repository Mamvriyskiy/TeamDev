package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
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
	auth.POST("/register", h.SignUp)

	

	// logger.Log("Info", "", "Create router", nil)

	return router
}