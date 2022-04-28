package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewGetUserProfilePhotos(userID int64) *GetUserProfilePhotos {
	return &GetUserProfilePhotos{
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotos struct {
	UserID int64  `json:"user_id"`
	Offset *int64 `json:"offset,omitempty"`
	Limit  *int64 `json:"limit,omitempty"`
}

func (gupp *GetUserProfilePhotos) Endpoint() string {
	return "getUserProfilePhotos"
}

func (gupp *GetUserProfilePhotos) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(gupp)
}

func (gupp *GetUserProfilePhotos) WithOffset(offset int64) {
	gupp.Offset = &offset
}

func (gupp *GetUserProfilePhotos) WithLimit(limit int64) {
	gupp.Limit = &limit
}
