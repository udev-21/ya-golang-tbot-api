package method

import (
	"github.com/udev-21/golang-tbot-api/types"
	"github.com/udev-21/golang-tbot-api/types/universal"
	"github.com/udev-21/golang-tbot-api/utils"
)

type AnswerInlineQuery struct {
	InlineQueryID     string                        `json:"inline_query_id"`
	Results           []universal.InlineQueryResult `json:"results"`
	CacheTime         *int64                        `json:"cache_time,omitempty"`
	IsPersonal        *bool                         `json:"is_personal,omitempty"`
	NextOffset        *string                       `json:"next_offset,omitempty"`
	SwitchPmText      *string                       `json:"switch_pm_text,omitempty"`
	SwitchPmParameter *string                       `json:"switch_pm_parameter,omitempty"`
}

func (apcq *AnswerInlineQuery) GetEndpoint() string {
	return "answerInlineQuery"
}

func (apcq *AnswerInlineQuery) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(apcq)
}

func (apcq *AnswerInlineQuery) UploadFiles() map[string]types.InputFile {
	return nil
}
