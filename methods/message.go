package methods

import (
	"encoding/json"

	"github.com/udev21/golang-tbot-api/types"
)

func NewSendMessage(text string) *Message {
	return &Message{
		Text: text,
	}
}

func (m Message) GetFieldsJsonPayload(chatID interface{}) ([]byte, error) {
	m.ChatID = chatID
	return json.Marshal(m)
}

func (m Message) ShouldSendFile() bool {
	return false
}

func (m Message) GetEndpoint() string {
	return "sendMessage"
}

type Message struct {
	ChatID                  interface{}           `json:"chat_id"`
	Text                    string                `json:"text"`
	ParseMode               *string               `json:"parse_mode,omitempty"`
	Entities                []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview   *bool                 `json:"disable_web_page_preview,omitempty"`
	DisableNotification     *bool                 `json:"disable_notification,omitempty"`
	ProtectContent          *bool                 `json:"protect_content,omitempty"`
	ReplyToMessageID        *int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithutReply *bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}

const (
	ParseModeMarkdownV2 string = "MarkdownV2"
	ParseModeHTML       string = "HTML"
	ParseModeMarkdown   string = "Markdown"
)

func (m *Message) WithParseModeHTML() *Message {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeHTML
	return m
}

func (m *Message) WithParseModeParseModeMarkdownV2() *Message {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeMarkdownV2
	return m
}

func (m *Message) WithParseModeParseModeMarkdown() *Message {
	m.ParseMode = new(string)
	*m.ParseMode = ParseModeMarkdown
	return m
}

// replyMarkup can be either InlineKeyboardMarkup
func (m *Message) WithReplyMarkup(replyMarkup interface{}) *Message {
	switch replyMarkup.(type) {
	case types.ReplyKeyboardMarkup, types.InlineKeyboardMarkup, types.ForceReply, types.ReplyKeyboardRemove:
		m.ReplyMarkup = replyMarkup
		return m
	default:
		panic("inappropriate type given")
	}
}
