package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"encoding/json"
	"bytes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const webhookURL = "https://acknowledge-chicago-vertical-lcd.trycloudflare.com/webhook"


func main() {
	bot, err := tgbotapi.NewBotAPI("7590824309:AAE_ocKJ0yIMpkqWIAiwBebtuEsfIo8o97A")
	if err != nil {
		log.Fatal(err)
	}

	// Set the webhook URL
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var update tgbotapi.Update
			err := json.NewDecoder(r.Body).Decode(&update)
			if err != nil {
				http.Error(w, "Ошибка раскодирования сообщения", http.StatusBadRequest)
				return
			}
			fmt.Println(update.Message.Text)

			log.Println("Пользователь:", update.Message.From, "Сообщение:", update.Message.Text)
			switch update.Message.Text {
			case "/start":
				url := "http://localhost:8000/auth/register"
				jsonData, err := json.Marshal(update.Message.From)
				if err != nil {
					log.Fatalf("Ошибка при преобразовании данных в JSON: %v", err)
				}

				// Отправка POST-запроса
				_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
				if err != nil {
					log.Fatalf("Ошибка при отправке запроса: %v", err)
				}
			}

			if update.Message != nil {
				handleMessage(bot, update.Message)
			}
		}
	})

	log.Println("Бот запущен...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	resp, err := http.Get("https://api64.ipify.org?format=text")
	if err != nil {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка получения IP"))
		return
	}
	defer resp.Body.Close()

	ip, _ := io.ReadAll(resp.Body)
	log.Println("Ваш IP: " +  string(ip))
	bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ваш IP: "+string(ip)))
}

// func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, ip string) {
// 	// menu := tgbotapi.NewInlineKeyboardMarkup(
// 	// 	tgbotapi.NewInlineKeyboardRow(
// 	// 		tgbotapi.NewInlineKeyboardButtonData("Профиль", "stats"),
// 	// 		tgbotapi.NewInlineKeyboardButtonData("Задачи", "settings"),
// 	// 	),
// 	// 	tgbotapi.NewInlineKeyboardRow(
// 	// 		tgbotapi.NewInlineKeyboardButtonURL("Купить монеты", "buy"),
// 	// 	),
// 	// )

// 	text := fmt.Sprintf("Ваш IP-адрес: %s", ip)
// 	reply := tgbotapi.NewMessage(message.Chat.ID, text)
// 	// reply.ReplyMarkup = menu

// 	log.Println("Отправка сообщения пользователю:", message.Chat.ID)

// 	_, err := bot.Send(reply)
// 	if err != nil {
// 		log.Println("Ошибка отправки сообщения:", err)
// 	}
// }

