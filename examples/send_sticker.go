package main

import (
	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/method/types"
)

func handle(ctx gtbotapi.Context) error {
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
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnText("sticker"), handle)

	bot.Start()
}
