package types

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
