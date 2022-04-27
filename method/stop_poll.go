package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewStopPoll(messageID int64) *StopPoll {
	return &StopPoll{
		MessageID: messageID,
	}
}

// https://core.telegram.org/bots/api#stoppoll
type StopPoll struct {
	MessageID int64 `json:"message_id"`
	myTypes.ReplyMarkuper
}

func (sp *StopPoll) Endpoint() string {
	return "stopPoll"
}

func (sp *StopPoll) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(sp)
}
