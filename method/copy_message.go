package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewCopyMessage(fromChatID string, messageID int64) *CopyMessage {
	return &CopyMessage{
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

// https://core.telegram.org/bots/api#copymessage
type CopyMessage struct {
	FromChatID string `json:"from_chat_id"`
	MessageID  int64  `json:"message_id"`

	myTypes.Captioner
	myTypes.ParseModer
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (a *CopyMessage) Endpoint() string {
	return "copyMessage"
}

func (a *CopyMessage) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(a)
}
