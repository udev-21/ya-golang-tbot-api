package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"

	"github.com/udev-21/golang-tbot-api/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendMessage(text string) *Message {
	return &Message{
		Text: text,
	}
}

func (m *Message) RawJsonPayload() (map[string]interface{}, error) {
	var err error

	// if m.ReplyMarkup != nil {
	// 	m.ReplyMarkup_ = make(map[string]interface{})
	// 	m.ReplyMarkup_, err = m.ReplyMarkup.RawJsonPayload()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	if err != nil {
		return nil, err
	}
	return utils.ConvertToMapStringInterface(m)
}

func (m *Message) UploadFiles() map[string]types.InputFile {
	return nil
}

func (m *Message) GetEndpoint() string {
	return "sendMessage"
}

type Message struct {
	// ChatID                string                `json:"chat_id"`
	Text                  string                `json:"text"`
	Entities              []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview *bool                 `json:"disable_web_page_preview,omitempty"`

	myTypes.ProtectContenter
	myTypes.ReplyMarkuper
	myTypes.ParseModer
	myTypes.ReplyToMessager
	myTypes.DisableNotificationer
}

func (p *Message) Endpoint() string {
	return "sendMessage"
}

func (p *Message) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(p)
}
