package filter

import golangtbotapi "github.com/udev-21/ya-golang-tbot-api"

func SuccessfulPayment(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.Message() != nil && ctx.Message().SuccessfulPayment != nil {
			next(ctx)
		}
		return nil
	}
}
