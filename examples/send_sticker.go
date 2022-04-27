package main

import (
	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/method/types"
	filter "github.com/udev-21/ya-golang-tbot-api/middleware"
)

func handle(ctx yagolangtbotapi.Context) error {
	sticker := types.NewFilePath("assets/sticker.tgs")
	msg := method.NewSendSticker(sticker)
	_, err := ctx.Send(ctx.Chat(), msg)
	if err != nil {
		panic(err)
	}
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnMessage("sticker"), handle)

	bot.Start()
}
