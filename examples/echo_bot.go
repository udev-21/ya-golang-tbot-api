package main

import (
	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
)

func handle(ctx gtbotapi.Context) error {
	text := "pong" // response text
	payload := method.NewSendMessage(text)

	payload.WithReplyToMessageID(ctx.Message().MessageID) //  reply to specific message

	reciver := ctx.Chat() // choose message reciever chat
	_, err := ctx.Send(reciver, payload)
	// apiResponse, err := ctx.Send(reciver, payload) // use this one if you need process api response

	return err
}

// replace value with your BOT TOKEN which gives you @botfather on telegram
const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	// if you send "ping" any chat this one will reply as "pong"
	bot.Handle(filter.OnText("ping"), handle)

	bot.Start()
}
