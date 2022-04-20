package methods

import "github.com/udev21/golang-tbot-api/types"

type AnswerWebAppQuery struct {
	WebAppQueryID string                  `json:"web_app_query_id"`
	Result        types.InlineQueryResult `json:"result"`
}
