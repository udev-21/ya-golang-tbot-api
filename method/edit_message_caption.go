package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type EditMessageCaption struct {
	MessageID       int64  `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`

	myTypes.Captioner
	myTypes.ParseModer
	myTypes.ReplyMarkuper
}

func (emc *EditMessageCaption) Endpoint() string {
	return "editMessageCaption"
}

func (emc *EditMessageCaption) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(emc)
}

func (emc *EditMessageCaption) WithMessageID(messageID int64) {
	emc.MessageID = messageID
}

func (emc *EditMessageCaption) WithInlineMessageID(inlineMessageID string) {
	emc.InlineMessageID = inlineMessageID
}
