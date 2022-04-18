package types

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}
