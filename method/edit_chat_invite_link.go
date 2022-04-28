package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewEditChatInviteLink(inviteLink string) *EditChatInviteLink {
	return &EditChatInviteLink{
		InviteLink: inviteLink,
	}
}

// https://core.telegram.org/bots/api#editchatinvitelink
type EditChatInviteLink struct {
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         *int64 `json:"expire_date,omitempty"`
	MemberLimit        *int64 `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

func (ecil *EditChatInviteLink) Endpoint() string {
	return "editChatInviteLink"
}

func (ecil *EditChatInviteLink) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(ecil)
}

func (ecil *EditChatInviteLink) WithName(name string) {
	ecil.Name = name
}

func (ecil *EditChatInviteLink) WithExpireDate(date int64) {
	ecil.ExpireDate = &date
}

func (ecil *EditChatInviteLink) WithLimit(limit int64) {
	ecil.MemberLimit = &limit
}

func (ecil *EditChatInviteLink) WithCreatesJoinRequest() {
	ecil.CreatesJoinRequest = true
}
