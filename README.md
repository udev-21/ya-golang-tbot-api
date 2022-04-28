# Telegram bot api client for golang
yet another Golang framework for Telegram Bot API client ([Telegram Passport](https://core.telegram.org/passport) included)
> I've seen, there's a lot good frameworks for this stuff, but without passport
> so that was my motivate ;)

```bash
go get -u github.com/udev-21/ya-golang-tbot-api
```


## Disclaimer
This project is my **PET project** and under [MIT license](https://opensource.org/licenses/MIT) :).


## Overview
This is an API client library for [Telegram Bot API](https://core.telegram.org/bots/api).
- [Simple Echo Bot](https://github.com/udev-21/ya-golang-tbot-api#simple-echo-bot)

Of course PR's and Forks are welcome :)

## Features
- [Filters](https://github.com/udev-21/ya-golang-tbot-api#filters)


## Simple Echo Bot
All examples you can find under [examples](https://github.com/udev-21/ya-golang-tbot-api/tree/main/examples) folder here's one of them:
```go
package main

import (
	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
	"github.com/udev-21/ya-golang-tbot-api/method"
)

func handle(ctx gtbotapi.Context) error {
	text := "pong" // response text
	payload := method.NewSendMessage(text)

	payload.WithReplyToMessageID(ctx.Message().MessageID) //  reply to specific message

	reciver := ctx.Chat() // choose message reciever chat
	_, err := ctx.Send(reciver, payload)
	// apiResponse, err := ctx.Send(reciver, payload) // use this one if you need process api response

	return err
}

// replace value with your BOT TOKEN which gives you @botfather on telegram 
const TOKEN = "BOT:TOKEN" 

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)

	// if you send "ping" any chat this one will reply as "pong"
	bot.Handle(filter.OnText("ping"), handle)

	bot.Start()
}

```


## Filters

Full list of filters you can find in [here](https://github.com/udev-21/ya-golang-tbot-api/tree/main/filter)

Here's some examples for using them:
```go
bot.Handle(filter.OnAnyText, yourHandlerFunc) // filter for: anyText message anywhere: private,group,supergroup

bot.Handle(filter.OnAnyText, yourHandlerFunc,filter.OnlySupergroup) // filter for: only supergroup and any text message

bot.Handle(filter.OnAnyText, yourHandlerFunc,filter.OnlySupergroupOrGroup) // filter for: only (supergroup or group) and any text message

bot.Handle(filter.OnAnyText, yourHandlerFunc, filter.OnlyPrivate) // filter for: only private and any text message

bot.Handle(filter.OnText("ping"), yourHandlerFunc, filter.OnlyPrivate) // filter for: only private and text message == "ping"
```


## Logging
