package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetChatTitle(title string) *SetChatTitle {
	return &SetChatTitle{
		Title: title,
	}
}

// https://core.telegram.org/bots/api#setchattitle
type SetChatTitle struct {
	Title string `json:"title"`
}

func (sct *SetChatTitle) Endpoint() string {
	return "setChatTitle"
}

func (sct *SetChatTitle) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(sct)
}
