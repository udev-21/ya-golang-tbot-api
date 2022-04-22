package types

type ChatMemberUpdated struct {
	Chat          Chat            `json:"chat"`
	Sender        User            `json:"from"`
	Date          int64           `json:"date"`
	OldChatMember ChatMember      `json:"old_chat_member"`
	NewChatMember ChatMember      `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}
