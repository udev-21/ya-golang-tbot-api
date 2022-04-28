package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

// https://core.telegram.org/bots/api#stopmessagelivelocation
func NewStopMessageLiveLocation() *StopMessageLiveLocation {
	return &StopMessageLiveLocation{}
}

type StopMessageLiveLocation struct {
	MessageID       string `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`

	myTypes.ReplyMarkuper
}

func (emll *StopMessageLiveLocation) Endpoint() string {
	return "stopMessageLiveLocation"
}

func (emll *StopMessageLiveLocation) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(emll)
}

func (emll *StopMessageLiveLocation) WithInlineMessageID(inlineMessageID string) {
	emll.InlineMessageID = inlineMessageID
}

func (emll *StopMessageLiveLocation) WithMessageID(messageID string) {
	emll.MessageID = messageID
}
