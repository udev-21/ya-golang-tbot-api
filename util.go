package yagolangtbotapi

import (
	"strconv"

	"github.com/udev-21/ya-golang-tbot-api/types"
)

func applyMiddlewares(handler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	return func(ctx Context) error {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler(ctx)
	}
}

func getChatID(chat *types.Chat) (string, error) {
	if chat == nil {
		return "", newError("chat is required")
	}
	var chatID string
	if chat.ID == 0 {
		if chat.Username == nil || len(*chat.Username) == 0 {
			return "", newError("chat is required")
		}
		chatID = *chat.Username
	} else {
		chatID = strconv.FormatInt(chat.ID, 10)
	}
	return chatID, nil
}
