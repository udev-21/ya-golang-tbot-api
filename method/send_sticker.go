package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendSticker(sticker myTypes.InputFile) *Sticker {
	return &Sticker{
		Sticker: sticker,
	}
}

type Sticker struct {
	Sticker myTypes.InputFile `json:"sticker"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (s *Sticker) Endpoint() string {
	return "sendSticker"
}

func (s *Sticker) Params() (myTypes.Params, error) {
	res, err := utils.ConvertToMapStringInterface(s)
	if err != nil {
		return nil, err
	}
	res["sticker"] = s.Sticker
	return res, nil
}

func (s *Sticker) Files() []myTypes.InputFile {
	if tmp, ok := s.Sticker.(myTypes.Uploadable); ok {
		tmp.SetField("sticker")
		return []myTypes.InputFile{
			tmp,
		}
	}
	return nil
}
