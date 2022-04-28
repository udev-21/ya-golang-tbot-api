package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewAnswerCallbackQuery(callbackQueryID string) *AnswerCallbackQuery {
	return &AnswerCallbackQuery{
		CallbackQueryID: callbackQueryID,
	}
}

// https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQuery struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int64  `json:"cache_time,omitempty"`
}

func (acq *AnswerCallbackQuery) Endpoint() string {
	return "answerCallbackQuery"
}

func (acq *AnswerCallbackQuery) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(acq)
}

func (acq *AnswerCallbackQuery) WithText(text string) {
	acq.Text = text
}

func (acq *AnswerCallbackQuery) WithShowAlert() {
	acq.ShowAlert = true
}

func (acq *AnswerCallbackQuery) WithUrl(url string) {
	acq.Url = url
}

func (acq *AnswerCallbackQuery) WithCacheTime(cacheTime int64) {
	acq.CacheTime = cacheTime
}
