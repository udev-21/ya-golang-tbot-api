package main

import (
	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func handle(ctx gtbotapi.Context) error {

	msg := method.NewSendMessage("message with keyboard")
	kMarkup := types.NewKeyboard(
		types.NewKeyboardRow(
			types.NewKeyboardButton("row1"),
		),
		types.NewKeyboardRow(
			types.NewKeyboardButton("row2 col1"),
			types.NewKeyboardButton("row2 col2"),
		),
		types.NewKeyboardRow(
			types.NewKeyboardButton("row3 col1"),
			types.NewKeyboardButton("row3 col2"),
			types.NewKeyboardButton("row3 col3"),
		),
	)

	kMarkup.WithResizeKeyboard()

	msg.WithReplyMarkup(kMarkup)

	_, err := ctx.Send(ctx.Chat(), msg)

	return err
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnAnyText, handle, filter.OnlyPrivate)

	bot.Start()
}
