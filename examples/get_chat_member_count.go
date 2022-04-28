package main

import (
	"fmt"

	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	filter "github.com/udev-21/ya-golang-tbot-api/middleware"
)

func handle(ctx yagolangtbotapi.Context) error {
	chatMember, err := ctx.GetChatMemberCount(ctx.Chat())
	if err != nil {
		panic(err)
	}

	msg := method.NewSendMessage(fmt.Sprintf("group count: %d", *chatMember))
	msg.WIthReplyToMessageID(ctx.Message().MessageID)
	_, err = ctx.Send(ctx.Chat(), msg)
	if err != nil {
		panic(err)
	}
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnMessage("count"), handle, filter.OnlySupergroupOrGroup)

	bot.Start()
}
