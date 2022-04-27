package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewEditMessageReplyMarkup() *EditMessageReplyMarkup {
	return &EditMessageReplyMarkup{}
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkup struct {
	MessageID       int64  `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`

	myTypes.ReplyMarkuper
}

func (emrm *EditMessageReplyMarkup) Endpoint() string {
	return "editMessageReplyMarkup"
}

func (emrm *EditMessageReplyMarkup) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(emrm)
}

func (emrm *EditMessageReplyMarkup) WithMessageID(messageID int64) {
	emrm.MessageID = messageID
}

func (emrm *EditMessageReplyMarkup) WithInlineMessageID(inlineMessageID string) {
	emrm.InlineMessageID = inlineMessageID
}

func (emrm *EditMessageReplyMarkup) WithReplyMarkup(replyMarkup types.InlineKeyboardMarkup) {
	emrm.ReplyMarkuper.ReplyMarkup = &replyMarkup
}
