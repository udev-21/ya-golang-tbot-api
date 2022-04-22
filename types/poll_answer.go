package types

type PollAnswer struct {
	PollID    string  `json:"poll_id"`
	Sender    User    `json:"user"`
	OptionIDs []int64 `json:"option_ids"`
}
