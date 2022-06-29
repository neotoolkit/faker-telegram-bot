package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
		args := strings.Split(update.Message.Text, " ")
		min := 0
		max := 100
		if len(args) >= 3 {
			convMin, err := strconv.Atoi(args[1])
			if nil == err {
				min = convMin
			}
			convMax, err := strconv.Atoi(args[2])
			if nil == err {
				max = convMax
			}
		}
		text = strconv.Itoa(f.Number(min, max))
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
