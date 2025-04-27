package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	fmt.Println("1")
	h.services.IUser.RegisterUser("daf")
}