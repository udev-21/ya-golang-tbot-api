package types

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	Sender          User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}
