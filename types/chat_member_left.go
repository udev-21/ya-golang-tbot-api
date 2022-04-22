package types

type ChatMemberLeft struct {
	Status string `json:"status"`
	Sender User   `json:"user"`
}
