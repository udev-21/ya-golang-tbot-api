package types

type ChatMember struct {
	Status              string  `json:"status"`
	Sender              User    `json:"user"`
	CanBeEdited         bool    `json:"can_be_edited"`
	IsAnonymous         bool    `json:"is_anonymous"`
	CanManageChat       bool    `json:"can_manage_chat"`
	CanDeleteMessages   bool    `json:"can_delete_messages"`
	CanManageVideoChats bool    `json:"can_manage_video_chats"`
	CanRestrictMembers  bool    `json:"can_restrict_members"`
	CanPromoteMembers   bool    `json:"can_promote_members"`
	CanChangeInfo       bool    `json:"can_change_info"`
	CanInviteUsers      bool    `json:"can_invite_users"`
	CanPostMessages     *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool   `json:"can_pin_messages,omitempty"`
	CustomTitle         *string `json:"custom_title,omitempty"`

	IsStillMember         bool  `json:"is_member"`
	CanSendMessages       bool  `json:"can_send_messages"`
	CanSendMediaMessages  bool  `json:"can_send_media_messages"`
	CanSendPolls          bool  `json:"can_send_polls"`
	CanSendOtherMessages  bool  `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews"`
	UntilDate             int64 `json:"until_date"`
}

const (
	ChatMemberTypeOwner         = "creator"
	ChatMemberTypeAdministrator = "administrator"
	ChatMemberTypeMember        = "member"
	ChatMemberTypeRestricted    = "restricted"
	ChatMemberTypeLeft          = "left"
	ChatMemberTypeBanned        = "kicked"
)

func (c ChatMember) IsOwner() bool {
	return c.Status == ChatMemberTypeOwner
}

func (c ChatMember) IsAdministrator() bool {
	return c.Status == ChatMemberTypeAdministrator
}

func (c ChatMember) IsMember() bool {
	return c.Status == ChatMemberTypeMember
}

func (c ChatMember) IsRestricted() bool {
	return c.Status == ChatMemberTypeRestricted
}

func (c ChatMember) IsLeft() bool {
	return c.Status == ChatMemberTypeLeft
}

func (c ChatMember) IsBanned() bool {
	return c.Status == ChatMemberTypeBanned
}
