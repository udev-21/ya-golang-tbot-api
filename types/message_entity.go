package types

type MessageEntity struct {
	Type     string  `json:"type"`
	Offset   int64   `json:"offset"`
	Length   int64   `json:"length"`
	Url      *string `json:"url,omitempty"`
	Sender   *User   `json:"user,omitempty"`
	Language *string `json:"language,omitempty"`
}
