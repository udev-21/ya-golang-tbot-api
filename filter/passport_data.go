package filter

import golangtbotapi "github.com/udev-21/ya-golang-tbot-api"

func OnPassportData(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Message() != nil && ctx.Message().PassportData != nil {
			next(ctx)
		}
		return nil
	}
}
