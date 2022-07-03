package main

import (
	"log"
	"net/http"
	"os"

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

	f := faker.NewFaker()

	switch update.Message.Command() {
	case "password":
		text = f.Internet().Password()
	case "username":
		text = f.Internet().Username()
	case "url":
		text = f.Internet().URL()
	case "domain":
		text = f.Internet().Domain()
	case "email":
		text = f.Internet().Email()
	case "number":
		text = number(&f, update.Message.Text)
	case "firstname":
		text = f.Person().FirstName()
	case "lastname":
		text = f.Person().LastName()
	case "name":
		text = f.Person().Name()
	case "color":
		text = f.Color().Color()
	case "hex":
		text = f.Color().Hex()
	case "uuid":
		text = f.UUID().V4()
	case "ipv4":
		text = f.Internet().IPv4()
	case "ipv6":
		text = f.Internet().IPv6()
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
