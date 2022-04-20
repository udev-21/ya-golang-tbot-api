package types

func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	return buttons
}

func NewInlineKeyboard(rows ...[]InlineKeyboardButton) [][]InlineKeyboardButton {
	return rows
}

func NewInlineKeyboardButton(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{
		Text: text,
	}
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	Url                          *string       `json:"url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo   `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl     `json:"login_url,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          *bool         `json:"pay,omitempty"`
}

func (ikb *InlineKeyboardButton) WithUrl(url string) *InlineKeyboardButton {
	ikb.Url = &url
	return ikb
}

func (ikb *InlineKeyboardButton) WithCallbackData(data string) *InlineKeyboardButton {
	ikb.CallbackData = &data
	return ikb
}

func (ikb *InlineKeyboardButton) WithWebApp(webApp WebAppInfo) *InlineKeyboardButton {
	ikb.WebApp = &webApp
	return ikb
}

func (ikb *InlineKeyboardButton) WithLoginUrl(loginUrl LoginUrl) *InlineKeyboardButton {
	ikb.LoginUrl = &loginUrl
	return ikb
}

func (ikb *InlineKeyboardButton) WithSwitchInlineQuery(input string) *InlineKeyboardButton {
	ikb.SwitchInlineQuery = &input
	return ikb
}

func (ikb *InlineKeyboardButton) WithSwitchInlineQueryCurrentChat(input string) *InlineKeyboardButton {
	ikb.SwitchInlineQueryCurrentChat = &input
	return ikb
}

func (ikb *InlineKeyboardButton) WithPay() *InlineKeyboardButton {
	pay := true
	ikb.Pay = &pay
	return ikb
}
