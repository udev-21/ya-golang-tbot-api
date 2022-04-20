package methods

import "github.com/udev21/golang-tbot-api/types"

type AnswerInlineQuery struct {
	InlineQueryID     string                    `json:"inline_query_id"`
	Results           []types.InlineQueryResult `json:"results"`
	CacheTime         *int64                    `json:"cache_time,omitempty"`
	IsPersonal        *bool                     `json:"is_personal,omitempty"`
	NextOffset        *string                   `json:"next_offset,omitempty"`
	SwitchPmText      *string                   `json:"switch_pm_text,omitempty"`
	SwitchPmParameter *string                   `json:"switch_pm_parameter,omitempty"`
}
