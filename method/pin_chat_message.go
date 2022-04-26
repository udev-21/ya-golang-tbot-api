package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewPinChatMessage(messageID int64) *PinChatMessage {
	return &PinChatMessage{
		MessageID: messageID,
	}
}

type PinChatMessage struct {
	MessageID int64 `json:"message_id"`
	myTypes.DisableNotificationer
}

func (pcm *PinChatMessage) Endpoint() string {
	return "pinChatMessage"
}

func (pcm *PinChatMessage) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(pcm)
}
