package types

type ChatMemberOwner struct {
	Status      string  `json:"status"`
	User        User    `json:"user"`
	IsAnonymous bool    `json:"is_anonymous"`
	CustomTitle *string `json:"custom_title,omitempty"`
}
