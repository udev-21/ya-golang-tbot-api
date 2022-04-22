package method

import (
	"time"

	"github.com/udev-21/golang-tbot-api/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

type GetUpdates struct {
	Offset         int64    `json:"offset,omitempty"`
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

func (gu *GetUpdates) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(gu)
}

func (gu GetUpdates) GetEndpoint() string {
	return "getUpdates"
}
func (gu GetUpdates) UploadFiles() map[string]types.InputFile {
	return nil
}
