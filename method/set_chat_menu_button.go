package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetChatMenuButton() *SetChatMenuButton {
	return &SetChatMenuButton{}
}

type SetChatMenuButton struct {
	ChatID     *int64             `json:"chat_id"`
	MenuButton *types.IMenuButton `json:"menu_button,omitempty"`
}

func (scmb *SetChatMenuButton) Endpoint() string {
	return "setChatMenuButton"
}

func (scmb *SetChatMenuButton) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(scmb)
}

func (scmb *SetChatMenuButton) WithMenuButton(menuButton types.IMenuButton) {
	scmb.MenuButton = &menuButton
}

func (scmb *SetChatMenuButton) WithChatID(chatID int64) {
	scmb.ChatID = &chatID
}
