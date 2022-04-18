package types

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  *bool                   `json:"request_contact,omitempty"`
	RequestLocation *bool                   `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo             `json:"web_app,omitempty"`
}
