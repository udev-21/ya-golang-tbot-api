package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewPromoteChatMember(userID int64) *PromoteChatMember {
	return &PromoteChatMember{
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#promotechatmember
type PromoteChatMember struct {
	UserID              int64 `json:"user_id"`
	IsAnonymous         bool  `json:"is_anonymous,omitempty"`
	CanManageChat       bool  `json:"can_manage_chat,omitempty"`
	CanPostMessages     bool  `json:"can_post_messages,omitempty"`
	CanEditMessages     bool  `json:"can_edit_messages,omitempty"`
	CanDeleteMessages   bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool  `json:"can_change_info,omitempty"`
	CanInviteUsers      bool  `json:"can_invite_users,omitempty"`
	CanPinMessages      bool  `json:"can_pin_messages,omitempty"`
}

func (pcm *PromoteChatMember) Endpoint() string {
	return "promoteChatMember"
}

func (pcm *PromoteChatMember) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(pcm)
}

func (cp *PromoteChatMember) WithIsAnonymous() {
	cp.IsAnonymous = true
}

func (cp *PromoteChatMember) WithCanPostMessages() {
	cp.CanPostMessages = true
}

func (cp *PromoteChatMember) WithCanEditMessages() {
	cp.CanEditMessages = true
}

func (cp *PromoteChatMember) WithCanDeleteMessages() {
	cp.CanDeleteMessages = true
}

func (cp *PromoteChatMember) WithCanManageVideoChats() {
	cp.CanManageVideoChats = true
}

func (cp *PromoteChatMember) WithCanRestrictMembers() {
	cp.CanRestrictMembers = true
}

func (cp *PromoteChatMember) WithCanPromoteMembers() {
	cp.CanPromoteMembers = true
}

func (cp *PromoteChatMember) WithCanChangeInfo() {
	cp.CanChangeInfo = true
}

func (cp *PromoteChatMember) WithCanManageChat() {
	cp.CanManageChat = true
}

func (cp *PromoteChatMember) WithCanInviteUsers() {
	cp.CanInviteUsers = true
}

func (cp *PromoteChatMember) WithCanPinMessages() {
	cp.CanPinMessages = true
}
