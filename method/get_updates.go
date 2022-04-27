package method

import (
	"time"

	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type GetUpdates struct {
	Offset         int64    `json:"offset"`
	Limit          *int64   `json:"limit,omitempty"`
	Timeout        *int64   `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (gu *GetUpdates) WithOffset(offset int64) *GetUpdates {
	gu.Offset = offset
	return gu
}

func (gu *GetUpdates) WithLimit(limit int64) *GetUpdates {
	gu.Limit = &limit
	return gu
}

func (gu *GetUpdates) WithTimeout(timeout time.Duration) *GetUpdates {
	val := int64(timeout / time.Second)
	gu.Timeout = &val
	return gu
}

func (gu *GetUpdates) WithAllowedUpdates(allowedUpdates []string) *GetUpdates {
	gu.AllowedUpdates = allowedUpdates
	return gu
}

func (gu *GetUpdates) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(gu)
}

func (gu GetUpdates) Endpoint() string {
	return "getUpdates"
}
