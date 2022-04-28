package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewAddStickerToSet(userId int64, name, emojis string) *AddStickerToSet {
	return &AddStickerToSet{
		StickerSetBase: StickerSetBase{
			UserID: userId,
			Emojis: emojis,
			Name:   name,
		},
	}
}

type AddStickerToSet struct {
	StickerSetBase
}

func (a *AddStickerToSet) Endpoint() string {
	return "addStickerToSet"
}

func (cnss *AddStickerToSet) Params() (myTypes.Params, error) {
	tmp, err := utils.ConvertToMapStringInterface(cnss)
	if err != nil {
		return nil, err
	}
	tmp["png_sticker"] = cnss.PngSticker
	tmp["tg_sticker"] = cnss.TgSticker
	tmp["webm_sticker"] = cnss.WebmSticker
	return tmp, nil
}

func (cnss *AddStickerToSet) Files() []myTypes.InputFile {
	return cnss.StickerSetBase.Files()
}
