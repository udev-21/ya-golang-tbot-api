package filter

import golangtbotapi "github.com/udev-21/golang-tbot-api"

func PreCheckout(next golangtbotapi.HandlerFunc) golangtbotapi.HandlerFunc {
	return func(ctx golangtbotapi.Context) error {
		if ctx.PreCheckoutQuery() != nil {
			next(ctx)
		}
		return nil
	}
}
