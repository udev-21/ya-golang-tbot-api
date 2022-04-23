package types

import "github.com/udev-21/golang-tbot-api/types"

type ReplyToMessager struct {
	ReplyToMessageID         *int64 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithutReply_ bool   `json:"allow_sending_without_reply,omitempty"`
}

func (m *ReplyToMessager) ReplyToMessage(msg *types.Message) {
	m.ReplyToMessageID = &msg.MessageID
}

func (m *ReplyToMessager) AllowSendingWithutReply() {
	m.AllowSendingWithutReply_ = true
}
