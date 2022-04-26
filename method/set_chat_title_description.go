package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSetChatDescription(description string) *SetChatDescription {
	return &SetChatDescription{
		Description: description,
	}
}

// https://core.telegram.org/bots/api#setchatdescription
type SetChatDescription struct {
	Description string `json:"description"`
}

func (scd *SetChatDescription) Endpoint() string {
	return "setChatDescription"
}

func (scd *SetChatDescription) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(scd)
}
