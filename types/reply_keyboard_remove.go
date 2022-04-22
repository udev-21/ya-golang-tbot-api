package types

import "github.com/udev-21/golang-tbot-api/utils"

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

func (rkm *ReplyKeyboardRemove) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(rkm)
}
