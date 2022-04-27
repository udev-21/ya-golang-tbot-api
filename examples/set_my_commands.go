package main

import (
	"log"

	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

const TOKEN = "BOT:TOKEN"

func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	commands := []types.BotCommand{
		types.BotCommand{
			Command:     "/ping",
			Description: "test is working",
		},
	}

	payload := method.NewSetMyCommands(commands)
	scope := types.NewBotCommandScopeAllPrivateChats()
	payload.WithScope(scope)

	log.Println(bot.SetMyCommands(payload))
}
