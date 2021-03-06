package types

type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard        *bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       *bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder *string             `json:"input_field_placeholder,omitempty"`
	Selective             *bool               `json:"selective,omitempty"`
}

func (rkm *ReplyKeyboardMarkup) WithResizeKeyboard() {
	rkm.ResizeKeyboard = new(bool)
	*rkm.ResizeKeyboard = true
}

func (rkm *ReplyKeyboardMarkup) WithOneTimeKeyboard() {
	rkm.OneTimeKeyboard = new(bool)
	*rkm.OneTimeKeyboard = true
}

func (rkm *ReplyKeyboardMarkup) WithInputFieldPlaceholder(placeholder string) {
	rkm.InputFieldPlaceholder = &placeholder
}

func (rkm *ReplyKeyboardMarkup) WithSelective() {
	rkm.Selective = new(bool)
	*rkm.Selective = true
}

// IsReplyMarkup - to satisfy ReplyMarkup interface
func (rkm *ReplyKeyboardMarkup) IsReplyMarkup() {}
