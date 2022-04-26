package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewUnpinChatMessage(messageID int64) *UnpinChatMessage {
	return &UnpinChatMessage{
		MessageID: messageID,
	}
}

type UnpinChatMessage struct {
	MessageID int64 `json:"message_id"`
}

func (upcm *UnpinChatMessage) Endpoint() string {
	return "unpinChatMessage"
}

func (ucm *UnpinChatMessage) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(ucm)
}
