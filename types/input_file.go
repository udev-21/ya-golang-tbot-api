package types

import (
	"fmt"
	"os"
)

// https://core.telegram.org/bots/api#inputfile
type InputFile interface {
	MustUpload() bool
	// should be either string or *os.File
	Value() interface{}
}

func NewInputFileLocal(file *os.File) (InputFile, error) {
	if file == nil {
		return nil, fmt.Errorf("given file invalid")
	}

	return &inputFileLocal{
		file: file,
	}, nil
}

func NewInputFileTelegram(fileID string) (InputFile, error) {
	return &inputFileTelegram{
		fileID: fileID,
	}, nil
}

func NewInputFileUrl(url string) (InputFile, error) {
	return &inputFileUrl{
		url: url,
	}, nil
}

type AlwaysUpload struct{}

func (mu *AlwaysUpload) MustUpload() bool {
	return true
}

type AlwaysNotUpload struct{}

func (mu *AlwaysNotUpload) MustUpload() bool {
	return false
}

type inputFileUrl struct {
	AlwaysNotUpload
	url string
}

func (ifu *inputFileUrl) Value() interface{} {
	return ifu.url
}

type inputFileTelegram struct {
	AlwaysNotUpload
	fileID string
}

func (ifu *inputFileTelegram) Value() interface{} {
	return ifu.fileID
}

type inputFileLocal struct {
	AlwaysUpload
	file *os.File
}

func (ifu *inputFileLocal) Value() interface{} {
	return ifu.file
}
