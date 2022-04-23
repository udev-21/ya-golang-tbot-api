package types

import (
	"encoding/json"
	"io"
	"os"
)

// RequestFileData represents the data to be used for a file.
type InputFile interface {
	// IsInputFile is for differentiate
	IsInputFile()
}

type Uploadable interface {
	InputFile

	// UploadData gets the file_attach_name and an `io.Reader` for the file to be uploaded. This
	// must only be called when the file needs to be uploaded.
	UploadData() (string, io.Reader, error)
}

// FileReader contains information about a reader to upload as a File.
type FileReader struct {
	Name   string
	Reader io.Reader
}

func (fr FileReader) MarshalJSON() ([]byte, error) {
	return json.Marshal(fr.Name)
}

func (fr FileReader) UploadData() (string, io.Reader, error) {
	return "attach://" + fr.Name, fr.Reader, nil
}

func (fr FileReader) IsInputFile() {}

// FilePath is a path to a local file.
type FilePath string

func (fp FilePath) MarshalJSON() ([]byte, error) {
	fileHandle, err := os.Open(string(fp))
	if err != nil {
		return nil, err
	}
	return json.Marshal(fileHandle.Name())
}

func (fp FilePath) UploadData() (string, io.Reader, error) {
	fileHandle, err := os.Open(string(fp))
	if err != nil {
		return "", nil, err
	}
	return "attach://" + fileHandle.Name(), fileHandle, err
}

func (fp FilePath) IsInputFile() {}

// FileURL is a URL to use as a file for a request.
type FileURL string

func (fu FileURL) IsInputFile() {}

// FileID is an ID of a file already uploaded to Telegram.
type FileID string

func (fi FileID) IsInputFile() {}
