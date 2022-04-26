package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
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

func (p *Audio) Endpoint() string {
	return "sendAudio"
}

func (p *Audio) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(p)
	if err != nil {
		return nil, err
	}
	params["audio"] = p.Audio
	return params, nil
}

func (p *Audio) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := p.Audio.(myTypes.Uploadable); ok {
		tmp.SetField("audio")
		res = append(res, tmp)
	}
	if tmp, ok := p.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}

func (p *Audio) WithPerformer(performer string) {
	p.Performer = &performer
}

func (p *Audio) WithTitle(title string) {
	p.Title = &title
}
