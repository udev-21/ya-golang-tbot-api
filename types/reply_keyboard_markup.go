package types

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"`
	Selective             *bool              `json:"selective,omitempty"`
}
