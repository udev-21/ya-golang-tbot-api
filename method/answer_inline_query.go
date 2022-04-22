package method

import "github.com/udev-21/golang-tbot-api/types/universal"

type AnswerInlineQuery struct {
	InlineQueryID     string                        `json:"inline_query_id"`
	Results           []universal.InlineQueryResult `json:"results"`
	CacheTime         *int64                        `json:"cache_time,omitempty"`
	IsPersonal        *bool                         `json:"is_personal,omitempty"`
	NextOffset        *string                       `json:"next_offset,omitempty"`
	SwitchPmText      *string                       `json:"switch_pm_text,omitempty"`
	SwitchPmParameter *string                       `json:"switch_pm_parameter,omitempty"`
}
