package main

import (
	"log"

	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
)

const TOKEN = "BOT:TOKEN"

func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	payload := method.NewDeleteMyCommands()
	// payload.WithScope()
	// payload.WithLanguageCode()

	if err := bot.DeleteMyCommands(payload); err != nil {
		panic(err)
	} else {
		log.Println("Ok")
	}
}
