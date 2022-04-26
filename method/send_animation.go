package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendAnimation(animation myTypes.InputFile) *Animation {
	return &Animation{
		Animation: animation,
	}
}

// https://core.telegram.org/bots/api#sendAnimation
type Animation struct {
	Animation myTypes.InputFile `json:"Animation"`
	Width     *int64            `json:"width,omitempty"`
	Height    *int64            `json:"height,omitempty"`

	myTypes.Durationer
	myTypes.Thumber
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (a *Animation) Endpoint() string {
	return "sendAnimation"
}

func (a *Animation) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(a)
	if err != nil {
		return nil, err
	}
	params["animation"] = a.Animation
	return params, nil
}

func (a *Animation) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := a.Animation.(myTypes.Uploadable); ok {
		tmp.SetField("animation")
		res = append(res, tmp)
	}
	if tmp, ok := a.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}

func (a *Animation) WithHeight(height int64) {
	a.Height = &height
}

func (a *Animation) WithWidth(width int64) {
	a.Width = &width
}
