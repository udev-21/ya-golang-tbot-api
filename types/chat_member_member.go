package types

type ChatMemberMember struct {
	Status string `json:"status"`
	Sender User   `json:"user"`
}
