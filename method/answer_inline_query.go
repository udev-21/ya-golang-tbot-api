package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewAnswerInlineQuery(inlineQueryID string, results []InlineQueryResult) *AnswerInlineQuery {
	return &AnswerInlineQuery{
		InlineQueryID: inlineQueryID,
		Results:       results,
	}
}

// https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQuery struct {
	InlineQueryID     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"results"`
	CacheTime         *int64              `json:"cache_time,omitempty"`
	IsPersonal        bool                `json:"is_personal,omitempty"`
	NextOffset        *string             `json:"next_offset,omitempty"`
	SwitchPmText      *string             `json:"switch_pm_text,omitempty"`
	SwitchPmParameter *string             `json:"switch_pm_parameter,omitempty"`
}

func (apcq *AnswerInlineQuery) Endpoint() string {
	return "answerInlineQuery"
}

func (apcq *AnswerInlineQuery) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(apcq)
}

func (apcq *AnswerInlineQuery) WithCacheTime(cacheTime int64) {
	apcq.CacheTime = &cacheTime
}

func (apcq *AnswerInlineQuery) WithIsPersonal() {
	apcq.IsPersonal = true
}

func (apcq *AnswerInlineQuery) WithNextOffset(nextOffset string) {
	apcq.NextOffset = &nextOffset
}

func (apcq *AnswerInlineQuery) WithSwitchPmText(text string) {
	apcq.SwitchPmText = &text
}

func (apcq *AnswerInlineQuery) WithSwitchPmParameter(parameter string) {
	apcq.SwitchPmParameter = &parameter
}
