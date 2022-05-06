package main

import (
	"encoding/json"
	"os"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func handle(ctx gtbotapi.Context) error {
	payload := method.NewInvoice(
		"Lenovo 3i Gaming",
		"Elevate your game and esports experience with a PC that'll stand tall in your gaming circle. Engineered with up to a cutting-edge 10th Gen Intel® Core™ i7 processor, discrete NVIDIA® graphics, up to a 120Hz FHD display, and an advanced gaming keyboard, the IdeaPad Gaming 3i laptop projects quiet, intimidating confidence. Enjoy clear visuals and tear-free gameplay on your path to domination.",
		"something",
		"PROVIDER:TOKEN", // provider token
		"USD",
		[]types.LabeledPrice{
			{
				Label:  "Laptop",
				Amount: 74999,
			},
		},
	)

	_, err := ctx.Send(ctx.Chat(), payload)
	// apiResponse, err := ctx.Send(reciver, payload) // use this one if you need process api response

	return err
}

// replace value with your BOT TOKEN which gives you @botfather on telegram
const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnText("payload"), handle)

	bot.Handle(filter.PreCheckout, preCheckoutHandler)

	bot.Start()
}

func preCheckoutHandler(ctx gtbotapi.Context) error {
	ans := method.NewAnswerPreCheckoutQuery(ctx.PreCheckoutQuery().ID, true)
	res, err := ctx.Send(ctx.Chat(), ans)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(os.Stdout).Encode(res)
	return err
}
