package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewVideoNote(videoNote myTypes.InputFile) *VideoNote {
	return &VideoNote{
		VideoNote: videoNote,
	}
}

// https://core.telegram.org/bots/api#sendvideonote
type VideoNote struct {
	VideoNote myTypes.InputFile `json:"video_note"`

	myTypes.Durationer
	myTypes.Thumber
	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkup
}

func (v *VideoNote) Endpoint() string {
	return "sendVideoNote"
}

func (v *VideoNote) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(v)
	if err != nil {
		return nil, err
	}
	params["video_note"] = v.VideoNote
	return params, nil
}

func (v *VideoNote) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := v.VideoNote.(myTypes.Uploadable); ok {
		tmp.SetField("video_note")
		res = append(res, tmp)
	}
	if tmp, ok := v.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}
