package types

type ChatMemberOwner struct {
	Status      string `json:"status"`
	Sender      User   `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
}
