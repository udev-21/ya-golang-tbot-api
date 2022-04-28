package method

type AnswerWebAppQuery struct {
	WebAppQueryID string            `json:"web_app_query_id"`
	Result        InlineQueryResult `json:"result"`
}
