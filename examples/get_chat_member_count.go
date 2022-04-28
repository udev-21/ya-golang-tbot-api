package main

import (
	"fmt"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
)

func handle(ctx gtbotapi.Context) error {
	chatMember, err := ctx.GetChatMemberCount(ctx.Chat())
	if err != nil {
		panic(err)
	}

	msg := method.NewSendMessage(fmt.Sprintf("group count: %d", *chatMember))
	msg.WithReplyToMessageID(ctx.Message().MessageID)
	_, err = ctx.Send(ctx.Chat(), msg)
	if err != nil {
		panic(err)
	}
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnText("count"), handle, filter.OnlySupergroupOrGroup)

	bot.Start()
}
