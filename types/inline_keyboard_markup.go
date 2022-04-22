package types

import "github.com/udev21/golang-tbot-api/utils"

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (ikb *InlineKeyboardMarkup) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(ikb)
}
