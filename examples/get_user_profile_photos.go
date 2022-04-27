package main

import (
	"encoding/json"
	"fmt"

	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	filter "github.com/udev-21/ya-golang-tbot-api/middleware"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func handle(ctx yagolangtbotapi.Context) error {

	req := method.NewGetUserProfilePhotos(ctx.Sender().ID)

	res, err := ctx.Send(nil, req)
	if err != nil {
		panic(err)
	} else if !res.OK {
		panic("photos not recieved")
	}

	var photos types.UserProfilePhotos
	if err = json.Unmarshal(res.Result, &photos); err != nil {
		panic(err)
	}

	msg := method.NewSendMessage(fmt.Sprintf("Total count of your pic`s: %d", photos.TotalCount))
	ctx.Send(ctx.Chat(), msg)
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.Message, handle, filter.OnlyPrivate)

	bot.Start()
}
