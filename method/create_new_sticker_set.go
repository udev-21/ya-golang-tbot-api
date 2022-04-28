package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type CreateNewStickerSet struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Emojis string `json:"emojis"`

	PngSticker    myTypes.InputFile   `json:"png_sticker,omitempty"`
	TgSticker     *myTypes.FilePath   `json:"tg_sticker,omitempty"`
	WebmSticker   *myTypes.FilePath   `json:"webm_sticker,omitempty"`
	ContainsMasks bool                `json:"contains_masks,omitempty"`
	MaskPosition  *types.MaskPosition `json:"mask_position,omitempty"`
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

func (cnss *CreateNewStickerSet) Files() []myTypes.InputFile {
	var files []myTypes.InputFile
	if tmp, ok := cnss.PngSticker.(myTypes.Uploadable); ok {
		tmp.SetField("png_sticker")
		files = append(files, tmp)
	}

	if cnss.TgSticker != nil {
		cnss.TgSticker.SetField("tg_sticker")
		files = append(files, cnss.PngSticker)
	}
	if cnss.WebmSticker != nil {
		cnss.WebmSticker.SetField("webm_sticker")
		files = append(files, cnss.WebmSticker)
	}
	return files
}

func (c *CreateNewStickerSet) WithPngSticker(pngSticker myTypes.InputFile) {
	c.PngSticker = pngSticker
}

func (c *CreateNewStickerSet) WithTgSticker(tgSticker myTypes.FilePath) {
	c.TgSticker = &tgSticker
}

func (c *CreateNewStickerSet) WithWebmSticker(webmSticker myTypes.FilePath) {
	c.WebmSticker = &webmSticker
}

func (c *CreateNewStickerSet) WithContainsMasks() {
	c.ContainsMasks = true
}

func (c *CreateNewStickerSet) WithMaskPosition(mp types.MaskPosition) {
	c.MaskPosition = &mp
}
