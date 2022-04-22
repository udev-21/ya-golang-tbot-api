package types

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	Sender     User            `json:"from"`
	Date       int64           `json:"date"`
	BIO        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
