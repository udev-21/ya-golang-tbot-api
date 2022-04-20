package types

import (
	"fmt"
	"os"
)

// https://core.telegram.org/bots/api#inputfile
type InputFile interface {
	MustUpload() bool
	// string or io.Reader
	Value() interface{}
}

type BaseInputFile struct{}

type AlwaysUpload struct{}

func (mu AlwaysUpload) MustUpload() bool {
	return true
}

type AlwaysNotUpload struct{}

func (mu AlwaysNotUpload) MustUpload() bool {
	return false
}

type InputFileUrl struct {
	BaseInputFile
	AlwaysNotUpload
	Url string
}

func (ifu InputFileUrl) Value() interface{} {
	return ifu.Url
}

type InputFileTelegram struct {
	BaseInputFile
	AlwaysNotUpload
	FileID string
}

func (ifu InputFileTelegram) Value() interface{} {
	return ifu.FileID
}

type InputFileLocal struct {
	BaseInputFile
	AlwaysUpload
	File *os.File
}

func (ifu InputFileLocal) Value() interface{} {
	return ifu.File
}

func NewInputFileLocal(file *os.File) (*InputFileLocal, error) {
	if file == nil {
		return nil, fmt.Errorf("given file invalid")
	}

	return &InputFileLocal{
		File: nil,
	}, nil
}
