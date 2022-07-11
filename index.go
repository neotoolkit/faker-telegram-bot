package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/neotoolkit/faker"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	update, err := bot.HandleUpdate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if nil == update.Message {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if len(update.Message.Command()) == 0 {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	var text string

	switch update.Message.Command() {
	case "password":
		text = faker.Password()
	case "username":
		text = faker.Username()
	case "url":
		text = faker.URL()
	case "domain":
		text = faker.Domain()
	case "email":
		text = faker.Email()
	case "number":
		text = number(update.Message.Text)
	case "firstname":
		text = faker.FirstName()
	case "lastname":
		text = faker.LastName()
	case "name":
		text = faker.Name()
	case "color":
		text = faker.Color()
	case "hex":
		text = faker.Hex()
	case "uuid":
		text = faker.UUID()
	case "ipv4":
		text = faker.IPv4()
	case "ipv6":
		text = faker.IPv6()
	case "bool":
		text = strconv.FormatBool(faker.Boolean())
	case "country":
		text = faker.Country()
	default:
		w.WriteHeader(http.StatusNoContent)

		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
