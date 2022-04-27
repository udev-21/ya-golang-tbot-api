package types

// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

func (cp *ChatPermissions) WithCanSendMessages() {
	cp.CanSendMessages = true
}
func (cp *ChatPermissions) WithCanSendMediaMessages() {
	cp.CanSendMediaMessages = true
}

func (cp *ChatPermissions) WithCanSendPolls() {
	cp.CanSendPolls = true
}

func (cp *ChatPermissions) WithCanSendOtherMessages() {
	cp.CanSendOtherMessages = true
}

func (cp *ChatPermissions) WithCanAddWebPagePreviews() {
	cp.CanAddWebPagePreviews = true
}

func (cp *ChatPermissions) WithCanChangeInfo() {
	cp.CanChangeInfo = true
}

func (cp *ChatPermissions) WithCanInviteUsers() {
	cp.CanInviteUsers = true
}

func (cp *ChatPermissions) WithCanPinMessages() {
	cp.CanPinMessages = true
}
