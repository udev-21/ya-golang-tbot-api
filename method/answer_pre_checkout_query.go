package method

import (
	"github.com/udev-21/golang-tbot-api/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string  `json:"pre_checkout_query_id"`
	OK                 bool    `json:"ok"`
	ErrorMessage       *string `json:"error_message"`
}

func (apcq *AnswerPreCheckoutQuery) GetEndpoint() string {
	return "answerPreCheckoutQuery"
}

func (apcq *AnswerPreCheckoutQuery) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(apcq)
}

func (apcq *AnswerPreCheckoutQuery) UploadFiles() map[string]types.InputFile {
	return nil
}

func (apcq *AnswerPreCheckoutQuery) WithErrorMessage(errorMsg string) {
	if apcq.ErrorMessage == nil {
		apcq.ErrorMessage = new(string)
	}
	*apcq.ErrorMessage = errorMsg
}
