package main

import (
	"log"

	yagolangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/method"
	filter "github.com/udev-21/ya-golang-tbot-api/middleware"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func handle(ctx yagolangtbotapi.Context) error {
	//first send new message with inlineKeyboard
	msg := method.NewSendMessage("new text message")
	kMarkup := types.NewInlineKeyboard(
		types.NewInlineKeyboardRow(
			types.NewInlineKeyboardButton("row1").WithCallbackData("callBackData1"),
		),
		types.NewInlineKeyboardRow(
			types.NewInlineKeyboardButton("row2 col1").WithCallbackData("callBackData2"),
			types.NewInlineKeyboardButton("row2 col2").WithCallbackData("callBackData3"),
		),
		types.NewInlineKeyboardRow(
			types.NewInlineKeyboardButton("row3 col1").WithCallbackData("callBackData4"),
			types.NewInlineKeyboardButton("row3 col2").WithCallbackData("callBackData5"),
			types.NewInlineKeyboardButton("row3 col3").WithCallbackData("callBackData6"),
		),
	)

	msg.WithReplyMarkup(kMarkup)

	response, err := ctx.Send(ctx.Chat(), msg)
	if err != nil {
		panic(err)
	} else if !response.OK {
		panic("not sended")
	}
	return nil
}

func answerCallbackQuery(ctx yagolangtbotapi.Context) error {
	log.Println("recieved data: ", ctx.CallbackQuery().Data)

	answerCQ := method.NewAnswerCallbackQuery(ctx.CallbackQuery().ID)
	answerCQ.WithText("this is answer text") // instead you can use WithUrl method, but anyway it's optional
	answerCQ.WithShowAlert()
	_, err := ctx.Send(ctx.Chat(), answerCQ)
	if err != nil {
		panic(err)
	}
	return nil
}

const TOKEN = "BOT:TOKEN"

func main() {
	bot := yagolangtbotapi.NewBotAPI(TOKEN)

	bot.Handle(filter.OnMessage("send"), handle)
	bot.Handle(filter.CallbackQuery, answerCallbackQuery)

	bot.Start()
}
