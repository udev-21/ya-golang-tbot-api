package types

// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  *bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
}

func (cp *ChatPermissions) WithCanSendMessages() {
	if cp.CanSendMessages != nil {
		cp.CanSendMessages = new(bool)
	}
	*cp.CanSendMessages = true
}
func (cp *ChatPermissions) WithCanSendMediaMessages() {
	if cp.CanSendMediaMessages != nil {
		cp.CanSendMediaMessages = new(bool)
	}
	*cp.CanSendMediaMessages = true
}

func (cp *ChatPermissions) WithCanSendPolls() {
	if cp.CanSendPolls != nil {
		cp.CanSendPolls = new(bool)
	}
	*cp.CanSendPolls = true
}

func (cp *ChatPermissions) WithCanSendOtherMessages() {
	if cp.CanSendOtherMessages != nil {
		cp.CanSendOtherMessages = new(bool)
	}
	*cp.CanSendOtherMessages = true
}

func (cp *ChatPermissions) WithCanAddWebPagePreviews() {
	if cp.CanAddWebPagePreviews != nil {
		cp.CanAddWebPagePreviews = new(bool)
	}
	*cp.CanAddWebPagePreviews = true
}

func (cp *ChatPermissions) WithCanChangeInfo() {
	if cp.CanChangeInfo != nil {
		cp.CanChangeInfo = new(bool)
	}
	*cp.CanChangeInfo = true
}

func (cp *ChatPermissions) WithCanInviteUsers() {
	if cp.CanInviteUsers != nil {
		cp.CanInviteUsers = new(bool)
	}
	*cp.CanInviteUsers = true
}

func (cp *ChatPermissions) WithCanPinMessages() {
	if cp.CanPinMessages != nil {
		cp.CanPinMessages = new(bool)
	}
	*cp.CanPinMessages = true
}
