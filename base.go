package golangtbotapi

import (
	"os"

	"github.com/udev21/golang-tbot-api/types"
)

type RawJsonPayloader interface {
	RawJsonPayload() (map[string]interface{}, error)
}

type MessagePayload interface {
	RawJsonPayloader
	GetEndpoint() string
	UploadFiles() map[string]types.InputFile
}

type HandlerFunc func(Context) error

type Middleware func(HandlerFunc) HandlerFunc

type UploadWithFiles interface {
	MessagePayload
	MustUploadFiles() map[string]*os.File
}
