package types

func NewKeyboardRow(buttons ...KeyboardButton) []KeyboardButton {
	return buttons
}

func NewKeyboard(rows ...[]KeyboardButton) [][]KeyboardButton {
	return rows
}

func NewKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{
		Text: text,
	}
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  *bool                   `json:"request_contact,omitempty"`
	RequestLocation *bool                   `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo             `json:"web_app,omitempty"`
}

func (kb *KeyboardButton) WithRequstContact() *KeyboardButton {
	kb.RequestContact = new(bool)
	*kb.RequestContact = true
	return kb
}

func (kb *KeyboardButton) WithRequstLocation() *KeyboardButton {
	kb.RequestLocation = new(bool)
	*kb.RequestLocation = true
	return kb
}
func (kb *KeyboardButton) WithWebApp(webApp WebAppInfo) *KeyboardButton {
	kb.WebApp = &webApp
	return kb
}
