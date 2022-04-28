package main

import (
	"encoding/json"
	"time"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func handle(ctx gtbotapi.Context) error {
	//first send new message to edit this one
	msg := method.NewSendMessage("new text message")
	response, err := ctx.Send(ctx.Chat(), msg)
	if err != nil {
		panic(err)
	} else if !response.OK {
		panic("not sended")
	}

	var sendedMessage types.Message

	if json.Unmarshal(response.Result, &sendedMessage) != nil {
		panic("can't unmarsal")
	}

	// wait for 2 seconds
	time.Sleep(2 * time.Second)

	edit := method.NewEditMessageText("new <b>EDITED</b> text message")
	edit.WithParseModeHTML()

	edit.WithMessageID(sendedMessage.MessageID)
	ctx.Send(ctx.Chat(), edit)
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnText("edit"), handle)

	bot.Start()
}
