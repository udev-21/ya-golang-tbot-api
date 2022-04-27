package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewEditMessageText(text string) *EditMessageText {
	return &EditMessageText{
		Text: text,
	}
}

type EditMessageText struct {
	Text                  string                `json:"text"`
	MessageID             int64                 `json:"message_id,omitempty"`
	InlineMessageID       string                `json:"inline_message_id,omitempty"`
	Entities              []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`

	myTypes.ParseModer
	myTypes.ReplyMarkuper
}

func (emt *EditMessageText) Endpoint() string {
	return "editMessageText"
}

func (emt *EditMessageText) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(emt)
}

func (emt *EditMessageText) WithMessageID(messageID int64) {
	emt.MessageID = messageID
}

func (emt *EditMessageText) WithInlineMessageID(inlineMessageID string) {
	emt.InlineMessageID = inlineMessageID
}

func (emt *EditMessageText) WithEntities(entities []types.MessageEntity) {
	emt.Entities = entities
}

func (emt *EditMessageText) WithDisableWebPagePreview() {
	emt.DisableWebPagePreview = true
}
