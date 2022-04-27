package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewForwardMessage(fromChatID string, messageID int64) *ForwardMessage {
	return &ForwardMessage{
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

// https://core.telegram.org/bots/api#forwardmessage
type ForwardMessage struct {
	FromChatID string `json:"from_chat_id"`
	MessageID  int64  `json:"message_id"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
}

func (fm *ForwardMessage) Endpoint() string {
	return "forwardMessage"
}

func (fm *ForwardMessage) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(fm)
}
