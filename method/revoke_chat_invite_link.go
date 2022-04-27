package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewRevokeChatInviteLink(inviteLink string) *RevokeChatInviteLink {
	return &RevokeChatInviteLink{
		InviteLink: inviteLink,
	}
}

// https://core.telegram.org/bots/api#revokechatinvitelink
type RevokeChatInviteLink struct {
	InviteLink string `json:"invite_link"`
}

func (ecil *RevokeChatInviteLink) Endpoint() string {
	return "revokeChatInviteLink"
}

func (ecil *RevokeChatInviteLink) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(ecil)
}
