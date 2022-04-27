package method

import (
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type GetUserProfilePhotos struct {
	UserID int64  `json:"user_id"`
	Offset *int64 `json:"offset,omitempty"`
	Limit  *int64 `json:"limit,omitempty"`
}

func (gupp *GetUserProfilePhotos) GetEndpoint() string {
	return "getUserProfilePhotos"
}

func (gupp *GetUserProfilePhotos) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(gupp)
}

func (gupp *GetUserProfilePhotos) UploadFiles() map[string]types.InputFile {
	return nil
}

func (gupp *GetUserProfilePhotos) WithOffset(offset int64) {
	if gupp.Offset == nil {
		gupp.Offset = new(int64)
	}
	gupp.Offset = &offset
}

func (gupp *GetUserProfilePhotos) WithLimit(limit int64) {
	if gupp.Limit == nil {
		gupp.Limit = new(int64)
	}
	gupp.Limit = &limit
}
