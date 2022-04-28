package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type StickerSetBase struct {
	UserID       int64               `json:"user_id"`
	Emojis       string              `json:"emojis"`
	Name         string              `json:"name"`
	PngSticker   myTypes.InputFile   `json:"png_sticker,omitempty"`
	TgSticker    *myTypes.FilePath   `json:"tg_sticker,omitempty"`
	WebmSticker  *myTypes.FilePath   `json:"webm_sticker,omitempty"`
	MaskPosition *types.MaskPosition `json:"mask_position,omitempty"`
}

func (c *StickerSetBase) WithPngSticker(pngSticker myTypes.InputFile) {
	c.PngSticker = pngSticker
}

func (c *StickerSetBase) WithTgSticker(tgSticker myTypes.FilePath) {
	c.TgSticker = &tgSticker
}

func (c *StickerSetBase) WithWebmSticker(webmSticker myTypes.FilePath) {
	c.WebmSticker = &webmSticker
}

func (c *StickerSetBase) WithMaskPosition(mp types.MaskPosition) {
	c.MaskPosition = &mp
}

func (cnss *StickerSetBase) Files() []myTypes.Uploadable {
	var files []myTypes.Uploadable
	if tmp, ok := cnss.PngSticker.(myTypes.Uploadable); ok {
		tmp.SetField("png_sticker")
		files = append(files, tmp)
	}

	if cnss.TgSticker != nil {
		cnss.TgSticker.SetField("tg_sticker")
		files = append(files, cnss.TgSticker)
	}

	if cnss.WebmSticker != nil {
		cnss.WebmSticker.SetField("webm_sticker")
		files = append(files, cnss.WebmSticker)
	}
	return files
}

func NewCreateNewStickerSet(userID int64, name, title, emojis string) *CreateNewStickerSet {
	return &CreateNewStickerSet{
		Title: title,
		StickerSetBase: StickerSetBase{
			UserID: userID,
			Name:   name,
			Emojis: emojis,
		},
	}
}

type CreateNewStickerSet struct {
	Title         string `json:"title"`
	ContainsMasks bool   `json:"contains_masks,omitempty"`
	StickerSetBase
}

func (cnss *CreateNewStickerSet) Endpoint() string {
	return "createNewStickerSet"
}

func (cnss *CreateNewStickerSet) Params() (myTypes.Params, error) {
	tmp, err := utils.ConvertToMapStringInterface(cnss)
	if err != nil {
		return nil, err
	}
	tmp["png_sticker"] = cnss.PngSticker
	tmp["tg_sticker"] = cnss.TgSticker
	tmp["webm_sticker"] = cnss.WebmSticker
	return tmp, nil
}

func (cnss *CreateNewStickerSet) Files() []myTypes.Uploadable {
	return cnss.StickerSetBase.Files()
}

func (c *CreateNewStickerSet) WithContainsMasks() {
	c.ContainsMasks = true
}
