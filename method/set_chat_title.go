package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSetChatTitle(title string) *SetChatTitle {
	return &SetChatTitle{
		Title: title,
	}
}

type SetChatTitle struct {
	Title string `json:"title"`
}

func (sct *SetChatTitle) Endpoint() string {
	return "setChatTitle"
}

func (sct *SetChatTitle) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(sct)
}
