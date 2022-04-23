package types

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// IsReplyMarkup - to satisfy ReplyMarkup interface
func (ikb *InlineKeyboardMarkup) IsReplyMarkup() {}
