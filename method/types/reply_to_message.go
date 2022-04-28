package types

type ReplyToMessager struct {
	ReplyToMessageID        *int64 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithutReply bool   `json:"allow_sending_without_reply,omitempty"`
}

func (m *ReplyToMessager) WithReplyToMessageID(id int64) {
	m.ReplyToMessageID = &id
}

func (m *ReplyToMessager) WithAllowSendingWithutReply() {
	m.AllowSendingWithutReply = true
}
