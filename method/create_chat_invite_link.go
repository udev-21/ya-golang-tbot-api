package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewCreateChatInviteLink() *CreateChatInviteLink {
	return &CreateChatInviteLink{}
}

// https://core.telegram.org/bots/api#createchatinvitelink
type CreateChatInviteLink struct {
	Name               string `json:"name,omitempty"`
	ExpireDate         *int64 `json:"expire_date,omitempty"`
	MemberLimit        *int64 `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

func (ccil *CreateChatInviteLink) Endpoint() string {
	return "createChatInviteLink"
}

func (ccil *CreateChatInviteLink) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(ccil)
}

func (ccil *CreateChatInviteLink) WithName(name string) {
	ccil.Name = name
}

func (ccil *CreateChatInviteLink) WithExpireDate(date int64) {
	ccil.ExpireDate = &date
}

func (ccil *CreateChatInviteLink) WithLimit(limit int64) {
	ccil.MemberLimit = &limit
}

func (ccil *CreateChatInviteLink) WithCreatesJoinRequest() {
	ccil.CreatesJoinRequest = true
}
