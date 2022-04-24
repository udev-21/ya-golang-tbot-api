package types

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// RequestFileData represents the data to be used for a file.
type InputFile interface {
	// IsInputFile is for differentiate
	IsInputFile()
}

type Uploadable interface {
	InputFile
	AttachName() string
	SetAttachName(string)

	UploadData() (io.Reader, error)
}

var _ Uploadable = &FilePath{}

type FilePath struct {
	Path       string
	attachName string
}

func NewFilePath(path string) *FilePath {
	return &FilePath{
		Path: path,
	}
}

func (fp FilePath) MarshalJSON() ([]byte, error) {
	if fp.attachName != "" {
		return json.Marshal("attach://" + fp.attachName)
	} else {
		return json.Marshal(filepath.Base(fp.Path))
	}
}

func (fp *FilePath) AttachName() string {
	// 'asdf'
	if fp.attachName != "" {
		return fp.attachName
	}
	return filepath.Base(fp.Path)
}

func (fp *FilePath) SetAttachName(name string) {
	fp.attachName = name
}

func (fp *FilePath) UploadData() (io.Reader, error) {
	fileHandle, err := os.Open(string(fp.Path))
	if err != nil {
		return nil, err
	}

	return fileHandle, nil
}

func (fp *FilePath) IsInputFile() {}

// FileURL is a URL to use as a file for a request.
type FileURL string

func (fu FileURL) IsInputFile() {}

// FileID is an ID of a file already uploaded to Telegram.
type FileID string

func (fi FileID) IsInputFile() {}
