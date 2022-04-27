package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type Photo struct {
	Photo           myTypes.InputFile     `json:"photo"`
	Caption         *string               `json:"caption,omitempty"`
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`

	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
	myTypes.ParseModer
	myTypes.DisableNotificationer
}

func (p *Photo) Endpoint() string {
	return "sendPhoto"
}

func (p *Photo) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(p)
	if err != nil {
		return nil, err
	}
	params["photo"] = p.Photo
	return params, nil
}

func (p *Photo) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := p.Photo.(myTypes.Uploadable); ok {
		tmp.SetField("photo")
		res = append(res, tmp)
	}
	return res
}
