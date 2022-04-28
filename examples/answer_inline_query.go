package main

import (
	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	filter "github.com/udev-21/ya-golang-tbot-api/middleware"
)

func handle(ctx yagolangtbotapi.Context) error {
	// results := []
	// ans := method.NewAnswerInlineQuery(ctx.InlineQuery().ID)

	return nil
}

const TOKEN = "BOT:TOKEN"

// FIRST OF ALL !!!
// you need to enable inline query status of your bot at @botfather
func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.InlineQuery, handle)

	bot.Start()
}
