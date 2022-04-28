package main

import (
	"encoding/json"
	"os"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
	mTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
)

func handle(ctx gtbotapi.Context) error {

	file1 := mTypes.NewFilePath("assets/image2.jpg")
	file1.SetCustomFileName("awesome-image.jpg")
	file2 := mTypes.NewFilePath("assets/image.jpg")

	thumb := mTypes.NewFilePath("assets/thumb2.jpg")

	media1 := mTypes.NewInputMediaDocument(file1)
	media1.WithCaption("caption1")

	media2 := mTypes.NewInputMediaDocument(file2)
	media2.WithCaption("caption2")

	media2.WithThumb(thumb)

	allMedias := []mTypes.SendMediaGroupable{
		media1,
		media2,
	}

	payload := method.NewSendMediaGroup(allMedias)
	res, err := ctx.Send(ctx.Chat(), payload)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(os.Stdout).Encode(res)
	return nil
}

// replace value with your BOT TOKEN which gives you @botfather on telegram
const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnText("media"), handle, filter.OnlyPrivate)

	bot.Start()
}
