package main

import (
	"log"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
)

func handle(ctx gtbotapi.Context) error {
	chatMember, err := ctx.GetChatMember(ctx.Chat(), ctx.Sender().ID)

	if err != nil {
		panic(err)
	}

	log.Println("isOwner: ", chatMember.IsOwner())
	// log.Println(*chatMember.Owner())
	log.Println("isAdmin: ", chatMember.IsAdministrator())
	// log.Println(*chatMember.Administrator())
	log.Println("isMember: ", chatMember.IsMember())
	// log.Println(*chatMember.Member())
	log.Println("isLeft: ", chatMember.IsLeft())
	// log.Println(*chatMember.Left())
	log.Println("isBanned: ", chatMember.IsBanned())
	// log.Println(*chatMember.Banned())
	log.Println("isRestricted: ", chatMember.IsRestricted())
	// log.Println(*chatMember.Restricted())

	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.AnyMessage, handle, filter.OnlySupergroupOrGroup)

	bot.Start()
}
