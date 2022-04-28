package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewUploadStickerFile(userID int64, pngSticker myTypes.FilePath) *UploadStickerFile {
	return &UploadStickerFile{
		UserID:     userID,
		PngSticker: pngSticker,
	}
}

type UploadStickerFile struct {
	UserID     int64            `json:"user_id"`
	PngSticker myTypes.FilePath `json:"png_sticker"`
}

func (usf *UploadStickerFile) Endpoint() string {
	return "uploadStickerFile"
}

func (usf *UploadStickerFile) Params() (myTypes.Params, error) {
	tmp, err := utils.ConvertToMapStringInterface(usf)
	if err != nil {
		return nil, err
	}
	tmp["png_sticker"] = usf.PngSticker
	return tmp, nil
}

func (usf *UploadStickerFile) Files() []myTypes.Uploadable {
	usf.PngSticker.SetField("png_sticker")
	return []myTypes.Uploadable{&usf.PngSticker}
}
