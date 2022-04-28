package main

import (
	"encoding/json"
	"os"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
)

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	payload := method.NewGetMyDefaultAdministratorRights()
	// payload.WithForChannels()
	rights, err := bot.GetMyDefaultAdministratorRights(payload)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(os.Stdout).Encode(rights)
}
