package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewEditMessageMedia() *EditMessageMedia {
	return &EditMessageMedia{}
}

// https://core.telegram.org/bots/api#editmessagemedia
type EditMessageMedia struct {
	Media           myTypes.InputMedia `json:"media"`
	MessageID       int64              `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`

	myTypes.ReplyMarkuper
}

func (emc *EditMessageMedia) Endpoint() string {
	return "editMessageMedia"
}

func (emc *EditMessageMedia) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(emc)
}

func (emc *EditMessageMedia) Files() []myTypes.Uploadable {
	return emc.Media.Files()
}

func (emc *EditMessageMedia) WithMessageID(messageID int64) {
	emc.MessageID = messageID
}

func (emc *EditMessageMedia) WithInlineMessageID(inlineMessageID string) {
	emc.InlineMessageID = inlineMessageID
}
