package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendAudio(audio myTypes.InputFile) *Audio {
	return &Audio{
		Audio: audio,
	}
}

// https://core.telegram.org/bots/api#sendaudio
type Audio struct {
	Audio     myTypes.InputFile `json:"audio"`
	Performer *string           `json:"performer,omitempty"`
	Title     *string           `json:"title,omitempty"`

	myTypes.Durationer
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.Thumber
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (a *Audio) Endpoint() string {
	return "sendAudio"
}

func (a *Audio) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(a)
	if err != nil {
		return nil, err
	}
	params["audio"] = a.Audio
	return params, nil
}

func (a *Audio) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := a.Audio.(myTypes.Uploadable); ok {
		tmp.SetField("audio")
		res = append(res, tmp)
	}
	if tmp, ok := a.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}

func (a *Audio) WithPerformer(performer string) {
	a.Performer = &performer
}

func (a *Audio) WithTitle(title string) {
	a.Title = &title
}
