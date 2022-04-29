package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewAnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool) *AnswerPreCheckoutQuery {
	return &AnswerPreCheckoutQuery{
		PreCheckoutQueryID: preCheckoutQueryID,
		OK:                 ok,
	}
}

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string  `json:"pre_checkout_query_id"`
	OK                 bool    `json:"ok"`
	ErrorMessage       *string `json:"error_message,omitempty"`
}

func (apcq *AnswerPreCheckoutQuery) Endpoint() string {
	return "answerPreCheckoutQuery"
}

func (apcq *AnswerPreCheckoutQuery) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(apcq)
}

func (apcq *AnswerPreCheckoutQuery) WithErrorMessage(errorMsg string) {
	if apcq.ErrorMessage == nil {
		apcq.ErrorMessage = new(string)
	}
	*apcq.ErrorMessage = errorMsg
}
