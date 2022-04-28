package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewAnswerWebAppQuery(webAppQueryID string, result InlineQueryResult) *AnswerWebAppQuery {
	return &AnswerWebAppQuery{
		WebAppQueryID: webAppQueryID,
		Result:        result,
	}
}

type AnswerWebAppQuery struct {
	WebAppQueryID string            `json:"web_app_query_id"`
	Result        InlineQueryResult `json:"result"`
}

func (awaq *AnswerWebAppQuery) Endpoint() string {
	return "answerWebAppQuery"
}

func (awaq *AnswerWebAppQuery) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(awaq)
}
