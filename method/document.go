package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type Document struct {
	Document        myTypes.InputFile     `json:"document"`
	Caption         *string               `json:"caption,omitempty"`
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`

	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`

	myTypes.Thumber
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
	myTypes.ParseModer
	myTypes.DisableNotificationer
}

func (p *Document) Endpoint() string {
	return "sendDocument"
}

func (p *Document) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(p)
	if err != nil {
		return nil, err
	}
	params["document"] = p.Document
	return params, nil
}

func (p *Document) Files() []myTypes.InputFile {
	var res []myTypes.InputFile
	if tmp, ok := p.Document.(myTypes.Uploadable); ok {
		tmp.SetField("document")
		res = append(res, tmp)
	}
	if tmp, ok := p.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}
