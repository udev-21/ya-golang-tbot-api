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
	CustomFileName() string
	FileName() string
	SetCustomFileName(string)
	SetField(string)
	Field() string

	UploadData() (io.Reader, error)
}

var _ Uploadable = &FilePath{}

type FilePath struct {
	Path           string
	customFileName string
	fieldName      string
}

func NewFilePath(path string) *FilePath {
	return &FilePath{
		Path: path,
	}
}

func (fp FilePath) MarshalJSON() ([]byte, error) {
	if fp.customFileName != "" {
		return json.Marshal("attach://" + fp.customFileName)
	} else {
		return json.Marshal(filepath.Base(fp.Path))
	}
}

func (fp *FilePath) CustomFileName() string {
	return fp.customFileName
}

func (fp *FilePath) FileName() string {
	if fp.customFileName != "" {
		return fp.customFileName
	} else {
		return filepath.Base(fp.Path)
	}
}

func (fp *FilePath) SetCustomFileName(name string) {
	fp.customFileName = name
}

func (fp *FilePath) SetField(field string) {
	fp.fieldName = field
}

func (fp *FilePath) Field() string {
	return fp.fieldName
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
type FileID FileURL

func (fi FileID) IsInputFile() {}
