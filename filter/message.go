package filter

import (
	golangtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func OnAnyMessage(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Message() != nil {
			next(ctx)
		}
		return nil
	}
}

func OnText(text string) golangtbotapi.Middleware {
	return func(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
		return OnAnyMessage(func(ctx golangtbotapi.Context) error {
			if ctx.Message().Text != nil && *ctx.Message().Text == text {
				next(ctx)
			}
			return nil
		})
	}
}

func OnAnyText(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return OnAnyMessage(func(ctx golangtbotapi.Context) error {
		if ctx.Message().Text != nil {
			next(ctx)
		}
		return nil
	})
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

func CallbackQuery(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.CallbackQuery() != nil {
			next(ctx)
		}
		return nil
	}
}

func OnlyGroup(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Chat() != nil && ctx.Chat().Type == types.ChatTypeGroup {
			next(ctx)
		}
		return nil
	}
}

func OnlyPrivate(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Chat() != nil && ctx.Chat().Type == types.ChatTypePrivate {
			next(ctx)
		}
		return nil
	}
}

func OnlySupergroup(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Chat() != nil && ctx.Chat().Type == types.ChatTypeSuperGroup {
			next(ctx)
		}
		return nil
	}
}

func OnlySupergroupOrGroup(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Chat() != nil && (ctx.Chat().Type == types.ChatTypeSuperGroup || ctx.Chat().Type == types.ChatTypeGroup) {
			next(ctx)
		}
		return nil
	}
}
