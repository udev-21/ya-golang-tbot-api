package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendDocument(document myTypes.InputFile) *Document {
	return &Document{
		Document: document,
	}
}

// https://core.telegram.org/bots/api#senddocument
type Document struct {
	Document                    myTypes.InputFile `json:"document"`
	DisableContentTypeDetection bool              `json:"disable_content_type_detection,omitempty"`

	myTypes.Captioner
	myTypes.Thumber
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
	myTypes.ParseModer
	myTypes.DisableNotificationer
}

func (d *Document) WithDisableContentTypeDetection() {
	d.DisableContentTypeDetection = true
}

func (d *Document) Endpoint() string {
	return "sendDocument"
}

func (d *Document) Params() (myTypes.Params, error) {
	var params myTypes.Params
	var err error
	params, err = utils.ConvertToMapStringInterface(d)
	if err != nil {
		return nil, err
	}
	params["document"] = d.Document
	return params, nil
}

func (d *Document) Files() []myTypes.Uploadable {
	var res []myTypes.Uploadable
	if tmp, ok := d.Document.(myTypes.Uploadable); ok {
		tmp.SetField("document")
		res = append(res, tmp)
	}
	if tmp, ok := d.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		res = append(res, tmp)
	}
	return res
}
