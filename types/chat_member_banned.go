package types

type ChatMemberBanned struct {
	Status    string `json:"status"`
	Sender    User   `json:"user"`
	UntilDate int64  `json:"until_date"`
}
