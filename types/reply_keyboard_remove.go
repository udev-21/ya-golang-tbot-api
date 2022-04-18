package types

type ReplyKeyboardRemove struct {
	//must be always true https://core.telegram.org/bots/api#replykeyboardremove
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}
