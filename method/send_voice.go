package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendVoice(voice myTypes.InputFile) *Voice {
	return &Voice{
		Voice: voice,
	}
}

// https://core.telegram.org/bots/api#sendvoice
type Voice struct {
	Voice myTypes.InputFile `json:"audio"`

	myTypes.Captioner
	myTypes.Durationer
	myTypes.ParseModer
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (v *Voice) Endpoint() string {
	return "sendVoice"
}

func (v *Voice) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(v)
	if err != nil {
		return nil, err
	}
	params["voice"] = v.Voice
	return params, nil
}

func (v *Voice) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := v.Voice.(myTypes.Uploadable); ok {
		tmp.SetField("voice")
		res = append(res, tmp)
	}
	return res
}
