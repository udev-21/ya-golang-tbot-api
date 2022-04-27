package types

import (
	"github.com/udev-21/golang-tbot-api/utils"
)

type ChatMember struct {
	Status              string  `json:"status"`
	User                User    `json:"user"`
	CanBeEdited         *bool   `json:"can_be_edited"`
	IsAnonymous         *bool   `json:"is_anonymous"`
	CanManageChat       *bool   `json:"can_manage_chat"`
	CanDeleteMessages   *bool   `json:"can_delete_messages"`
	CanManageVideoChats *bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  *bool   `json:"can_restrict_members"`
	CanPromoteMembers   *bool   `json:"can_promote_members"`
	CanChangeInfo       *bool   `json:"can_change_info"`
	CanInviteUsers      *bool   `json:"can_invite_users"`
	CanPostMessages     *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool   `json:"can_pin_messages,omitempty"`
	CustomTitle         *string `json:"custom_title,omitempty"`

	IsStillMember         *bool  `json:"is_member"`
	CanSendMessages       *bool  `json:"can_send_messages"`
	CanSendMediaMessages  *bool  `json:"can_send_media_messages"`
	CanSendPolls          *bool  `json:"can_send_polls"`
	CanSendOtherMessages  *bool  `json:"can_send_other_messages"`
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews"`
	UntilDate             *int64 `json:"until_date"`
}

const (
	ChatMemberTypeOwner         = "creator"
	ChatMemberTypeAdministrator = "administrator"
	ChatMemberTypeMember        = "member"
	ChatMemberTypeRestricted    = "restricted"
	ChatMemberTypeLeft          = "left"
	ChatMemberTypeBanned        = "kicked"
)

func (cm *ChatMember) IsOwner() bool {
	return cm.Status == ChatMemberTypeOwner
}

func (cm *ChatMember) Owner() (owner *ChatMemberOwner) {
	if cm.IsOwner() {
		//check required fields according to https://core.telegram.org/bots/api#chatmemberowner
		if !utils.NotNil(cm.IsAnonymous) {
			return
		}
		owner = &ChatMemberOwner{
			Status:      cm.Status,
			User:        cm.User,
			IsAnonymous: *cm.IsAnonymous,
			CustomTitle: cm.CustomTitle,
		}
	}
	return
}

func (cm *ChatMember) IsAdministrator() bool {
	return cm.Status == ChatMemberTypeAdministrator
}

func (cm *ChatMember) Administrator() (cma *ChatMemberAdministrator) {
	if cm.IsAdministrator() {
		//check required fields according to https://core.telegram.org/bots/api#chatmemberadministrator
		if !utils.NotNil(
			cm.CanBeEdited,
			cm.IsAnonymous,
			cm.CanManageChat,
			cm.CanDeleteMessages,
			cm.CanManageVideoChats,
			cm.CanRestrictMembers,
			cm.CanPromoteMembers,
			cm.CanChangeInfo,
			cm.CanInviteUsers,
		) {
			return
		}
		cma = &ChatMemberAdministrator{
			Status:              cm.Status,
			User:                cm.User,
			CanBeEdited:         *cm.CanBeEdited,
			CanManageChat:       *cm.CanManageChat,
			CanDeleteMessages:   *cm.CanDeleteMessages,
			CanManageVideoChats: *cm.CanManageVideoChats,
			CanRestrictMembers:  *cm.CanRestrictMembers,
			CanPromoteMembers:   *cm.CanPromoteMembers,
			CanChangeInfo:       *cm.CanChangeInfo,
			CanInviteUsers:      *cm.CanInviteUsers,
			CanPostMessages:     cm.CanPostMessages,
			CanEditMessages:     cm.CanEditMessages,
			CanPinMessages:      cm.CanPinMessages,
			CustomTitle:         cm.CustomTitle,
		}
	}
	return
}

func (cm *ChatMember) IsMember() bool {
	return cm.Status == ChatMemberTypeMember
}

func (cm *ChatMember) Member() (member *ChatMemberMember) {
	if cm.IsMember() {
		return &ChatMemberMember{
			Status: cm.Status,
			User:   cm.User,
		}
	}
	return
}

func (cm *ChatMember) IsRestricted() bool {
	return cm.Status == ChatMemberTypeRestricted
}

func (cm *ChatMember) Restricted() (r *ChatMemberRestricted) {
	if cm.IsRestricted() {
		//check required fields according to https://core.telegram.org/bots/api#chatmemberrestricted
		if !utils.NotNil(
			cm.IsStillMember,
			cm.CanChangeInfo,
			cm.CanInviteUsers,
			cm.CanPinMessages,
			cm.CanSendMessages,
			cm.CanSendMediaMessages,
			cm.CanSendPolls,
			cm.CanSendOtherMessages,
			cm.CanAddWebPagePreviews,
			cm.UntilDate,
		) {
			return
		}
		r = &ChatMemberRestricted{
			Status:                cm.Status,
			User:                  cm.User,
			IsMember:              *cm.IsStillMember,
			CanChangeInfo:         *cm.CanChangeInfo,
			CanInviteUsers:        *cm.CanInviteUsers,
			CanPinMessages:        *cm.CanPinMessages,
			CanSendMessages:       *cm.CanSendMessages,
			CanSendMediaMessages:  *cm.CanSendMediaMessages,
			CanSendPolls:          *cm.CanSendPolls,
			CanSendOtherMessages:  *cm.CanSendOtherMessages,
			CanAddWebPagePreviews: *cm.CanAddWebPagePreviews,
			UntilDate:             *cm.UntilDate,
		}
	}
	return
}

func (cm *ChatMember) IsLeft() bool {
	return cm.Status == ChatMemberTypeLeft
}

func (cm *ChatMember) Left() (cml *ChatMemberLeft) {
	if cm.IsLeft() {
		cml = &ChatMemberLeft{
			Status: cm.Status,
			User:   cm.User,
		}
	}
	return
}

func (cm *ChatMember) IsBanned() bool {
	return cm.Status == ChatMemberTypeBanned
}

func (cm *ChatMember) Banned() (b *ChatMemberBanned) {
	if cm.IsBanned() {
		//check required fields according to https://core.telegram.org/bots/api#chatmemberbanned
		if !utils.NotNil(cm.UntilDate) {
			return
		}
		b = &ChatMemberBanned{
			Status:    cm.Status,
			User:      cm.User,
			UntilDate: *cm.UntilDate,
		}
	}
	return
}
