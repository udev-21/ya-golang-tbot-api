package types

func NewReplyKeyboardRemove() *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
}

type ReplyKeyboardRemove struct {
	//must be always true https://core.telegram.org/bots/api#replykeyboardremove
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

func (rkr *ReplyKeyboardRemove) WithSelective(selective bool) *ReplyKeyboardRemove {
	rkr.Selective = &selective
	return rkr
}

// IsReplyMarkup - to satisfy ReplyMarkup interface
func (rkm *ReplyKeyboardRemove) IsReplyMarkup() {}
