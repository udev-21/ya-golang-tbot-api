package main

import (
	"log"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	commands := []types.BotCommand{
		types.BotCommand{
			Command:     "/ping",
			Description: "test is working",
		},
	}

	payload := method.NewSetMyCommands(commands)
	// payload.WithScope(scope)

	log.Println(bot.SetMyCommands(payload))
}
