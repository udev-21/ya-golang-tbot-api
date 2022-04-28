package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"

	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendMessage(text string) *Message {
	return &Message{
		Text: text,
	}
}

// https://core.telegram.org/bots/api#sendmessage
type Message struct {
	Text                  string                `json:"text"`
	Entities              []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`

	myTypes.ParseModer
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (p *Message) Endpoint() string {
	return "sendMessage"
}

func (p *Message) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(p)
}

func (p *Message) WithEntities(entieies []types.MessageEntity) {
	p.Entities = entieies
}

func (p *Message) WithDisableWebPagePreview() {
	p.DisableWebPagePreview = true
}
