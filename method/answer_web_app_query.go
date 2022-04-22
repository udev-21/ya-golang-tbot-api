package method

import "github.com/udev21/golang-tbot-api/types/universal"

type AnswerWebAppQuery struct {
	WebAppQueryID string                      `json:"web_app_query_id"`
	Result        universal.InlineQueryResult `json:"result"`
}
