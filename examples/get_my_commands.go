package main

import (
	"encoding/json"
	"os"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
)

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	payload := method.NewGetMyCommands()
	// payload.WithScope()
	// payload.WithLanguageCode()

	myCommands, err := bot.GetMyCommands(payload)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(os.Stdout).Encode(myCommands)
}
