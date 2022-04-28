package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetStickerSetThumb(userID int64, name string) *SetStickerSetThumb {
	return &SetStickerSetThumb{
		Name:   name,
		UserID: userID,
	}
}

type SetStickerSetThumb struct {
	Name   string            `json:"name"`
	UserID int64             `json:"user_id"`
	Thumb  myTypes.InputFile `json:"thumb"`
}

func (ssst *SetStickerSetThumb) Endpoint() string {
	return "setStickerSetThumb"
}
func (ssst *SetStickerSetThumb) Params() (myTypes.Params, error) {
	old, err := utils.ConvertToMapStringInterface(ssst)
	if err != nil {
		return nil, err
	}

	if ssst.Thumb != nil {
		old["thumb"] = ssst.Thumb
	}

	return old, nil
}

func (ssst *SetStickerSetThumb) Files() []myTypes.Uploadable {
	if tmp, ok := ssst.Thumb.(myTypes.Uploadable); ok {
		tmp.SetField("thumb")
		return []myTypes.Uploadable{tmp}
	}
	return nil
}

func (ssst *SetStickerSetThumb) WithThumb(thumb myTypes.InputFile) {
	ssst.Thumb = thumb
}
