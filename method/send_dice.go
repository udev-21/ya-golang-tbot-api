package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendDice() *Dice {
	return &Dice{}
}

// https://core.telegram.org/bots/api#senddice
type Dice struct {
	Emoji string `json:"emoji,omitempty"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (d *Dice) Endpoint() string {
	return "sendDice"
}

func (d *Dice) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(d)
}

func (d *Dice) WithEmoji(emoji string) {
	d.Emoji = emoji
}
