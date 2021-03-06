package main

import (
	"encoding/json"
	"os"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
)

const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)
	// to get default menuButton or replace nil with specific chat and result will be for that chat
	button, err := bot.GetChatMenuButton(nil)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(os.Stdout).Encode(button)
}
