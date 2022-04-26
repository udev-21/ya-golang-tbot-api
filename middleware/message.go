package filter

import (
	golangtbotapi "github.com/udev-21/golang-tbot-api"
)

func Message(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Message() != nil {
			next(ctx)
		}
		return nil
	}
}

func OnMessage(text string) golangtbotapi.Middleware {
	return func(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
		return Message(func(ctx golangtbotapi.Context) error {
			if ctx.Message().Text != nil && *ctx.Message().Text == text {
				next(ctx)
			}
			return nil
		})
	}
}

func EditedMessage(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.EditedMessage() != nil {
			next(ctx)
		}
		return nil
	}
}

func InlineQuery(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.InlineQuery() != nil {
			next(ctx)
		}
		return nil
	}
}
