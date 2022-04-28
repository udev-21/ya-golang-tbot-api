package main

import (
	"encoding/json"
	"log"
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
	if ctx.DeleteMessage(ctx.Chat(), sendedMessage.MessageID) != nil {
		panic("not deleted")
	}
	log.Println("message with id(%d) was deleted", sendedMessage.MessageID)
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnText("delete"), handle)

	bot.Start()
}
