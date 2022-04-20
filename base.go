package golangtbotapi

import (
	"github.com/udev21/golang-tbot-api/methods"
	"github.com/udev21/golang-tbot-api/types"
)

type MessagePayload interface {
	GetEndpoint() string
	GetFieldsJsonPayload(revicer interface{}) ([]byte, error)
}

var _ MessagePayload = (*methods.Message)(nil)

type UploadFiler interface {
	MessagePayload
	UploadFiles() types.InputFile
}
