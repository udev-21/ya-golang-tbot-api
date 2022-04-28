package main

import (
	"math/rand"

	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	filter "github.com/udev-21/ya-golang-tbot-api/middleware"
)

func someRandomGenerator() string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func handle(ctx yagolangtbotapi.Context) error {
	var results []method.InlineQueryResult
	for i := 0; i < 15; i++ {
		result := method.NewInlineQueryResultPhoto(
			someRandomGenerator(),
			"https://image.shutterstock.com/z/stock-vector-example-stamp-vector-template-illustration-design-vector-eps-1362869099.jpg",
			"https://image.shutterstock.com/image-vector/example-stamp-vector-template-illustration-260nw-1362869099.jpg",
		)
		results = append(results, result)
	}
	ans := method.NewAnswerInlineQuery(ctx.InlineQuery().ID, results)

	_, err := ctx.Send(nil, ans)
	if err != nil {
		panic(err)
	}

	return nil
}

const TOKEN = "BOT:TOKEN"

// FIRST OF ALL !!!
// you need to enable inline query status of your bot at @botfather
func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.InlineQuery, handle)

	bot.Start()
}
