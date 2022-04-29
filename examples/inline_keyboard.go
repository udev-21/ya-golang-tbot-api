package main

import (
	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func handle(ctx gtbotapi.Context) error {
	//first send new message with inlineKeyboard
	msg := method.NewSendMessage("new text message")
	kMarkup := types.NewInlineKeyboard(
		types.NewInlineKeyboardRow(
			types.NewInlineKeyboardButton("row1").WithCallbackData("callBackData1"),
		),
		types.NewInlineKeyboardRow(
			types.NewInlineKeyboardButton("row2 col1").WithCallbackData("callBackData2"),
			types.NewInlineKeyboardButton("row2 col2").WithCallbackData("callBackData3"),
		),
		types.NewInlineKeyboardRow(
			types.NewInlineKeyboardButton("row3 col1").WithCallbackData("callBackData4"),
			types.NewInlineKeyboardButton("row3 col2").WithCallbackData("callBackData5"),
			types.NewInlineKeyboardButton("row3 col3").WithCallbackData("callBackData6"),
		),
	)

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
