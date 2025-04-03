package main

import (
	// "fmt"
	"log"
	"net/http"
	"io"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const webhookURL = "https://implications-attribute-vulnerable-jul.trycloudflare.com/webhook"

func getClientIP(r *http.Request) string {
	// Попробуем получить IP через заголовки (если есть прокси)
	if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	return r.RemoteAddr // Обычный способ
}

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
			clientIP := getClientIP(r)
			log.Println("IP отправителя запроса:", clientIP)
			var update tgbotapi.Update
			err := json.NewDecoder(r.Body).Decode(&update)
			if err != nil {
				http.Error(w, "Ошибка раскодирования сообщения", http.StatusBadRequest)
				return
			}

			log.Println("Пользователь:", update.Message.From, "Сообщение:", update.Message.Text)
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
	log.Println("Ваш IP: "+string(ip))
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

