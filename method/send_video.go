package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendVideo(video myTypes.InputFile) *Video {
	return &Video{
		Video: video,
	}
}

// https://core.telegram.org/bots/api#sendvideo
type Video struct {
	Video             myTypes.InputFile `json:"video"`
	Width             *int64            `json:"width,omitempty"`
	Height            *int64            `json:"height,omitempty"`
	SupportsStreaming bool              `json:"supports_streaming,omitempty"`

	myTypes.Durationer
	myTypes.Thumber
	myTypes.Captioner
	myTypes.ParseModer
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (v *Video) Endpoint() string {
	return "sendVideo"
}

func (v *Video) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(v)
	if err != nil {
		return nil, err
	}
	params["video"] = v.Video
	return params, nil
}

func (v *Video) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := v.Video.(myTypes.Uploadable); ok {
		tmp.SetField("video")
		res = append(res, tmp)
	}
	if tmp, ok := v.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}

func (v *Video) WithHeight(height int64) {
	v.Height = &height
}

func (v *Video) WithWidth(width int64) {
	v.Width = &width
}

func (v *Video) WithSupportsStreaming() {
	v.SupportsStreaming = true
}
