package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var input tgbotapi.Update

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		log.Fatal("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	h.services.IUser.RegisterUser(input.Message.From.ID)
}

func (h *Handler) ProfileUser(c *gin.Context) {
	var input tgbotapi.Update
	fmt.Println("aaaaaa")

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		log.Fatal("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	fmt.Println(h.services.IUser.ProfileUser(input.Message.From.ID))
}

func (h *Handler) AddSocialUser(c *gin.Context) {
	var input tgbotapi.Update
	fmt.Println("bbbbb")

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		log.Fatal("Error", "c.BindJSON()", "Error bind json:", err, "")
		return
	}

	fmt.Println(h.services.IUser.AddSocialUser(input.Message.From.ID, input.Message.Text))
}
