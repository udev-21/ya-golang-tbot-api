package types

import (
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type InputMedia interface {
	Filer
	IsInputMedia()
}

type BaseInputMedia struct {
	Type  string    `json:"type"`
	Media InputFile `json:"media"`

	ParseModer
	Captioner
}

func (b *BaseInputMedia) IsInputMedia() {}

func (b *BaseInputMedia) Files() []InputFile {
	return []InputFile{b.Media}
}
func (b *BaseInputMedia) WithMedia(file InputFile) {
	b.Media = file
}

func (b *BaseInputMedia) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(b)
}

type InputMediaPhoto struct {
	BaseInputMedia
}

type InputMediaVideo struct {
	BaseInputMedia
	Thumb             InputFile `json:"thumb,omitempty"`
	Width             *int64    `json:"width,omitempty"`
	Height            *int64    `json:"height,omitempty"`
	SupportsStreaming bool      `json:"supports_streaming,omitempty"`
	Durationer
}

func (bmv *InputMediaVideo) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(bmv)
}

func (imv *InputMediaVideo) Files() []InputFile {
	files := imv.BaseInputMedia.Files()
	if imv.Thumb != nil {
		files = append(files, imv.Thumb)
	}
	return files
}

type InputMediaAnimation struct {
	BaseInputMedia
	Thumb  InputFile `json:"thumb,omitempty"`
	Width  *int64    `json:"width,omitempty"`
	Height *int64    `json:"height,omitempty"`
	Durationer
}

func (ima *InputMediaAnimation) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(ima)
}

func (ima *InputMediaAnimation) Files() []InputFile {
	files := ima.BaseInputMedia.Files()
	if ima.Thumb != nil {
		files = append(files, ima.Thumb)
	}
	return files
}

type InputMediaAudio struct {
	BaseInputMedia
	Durationer
	Thumb     InputFile `json:"thumb,omitempty"`
	Performer *string   `json:"performer,omitempty"`
	Title     *string   `json:"title,omitempty"`
}

func (ima *InputMediaAudio) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(ima)
}

func (ima *InputMediaAudio) Files() []InputFile {
	files := ima.BaseInputMedia.Files()
	if ima.Thumb != nil {
		files = append(files, ima.Thumb)
	}
	return files
}

type InputMediaDocument struct {
	BaseInputMedia
	Thumb                       InputFile `json:"thumb,omitempty"`
	DisableContentTypeDetection bool      `json:"disable_content_type_detection,omitempty"`
}

func (document *InputMediaDocument) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(document)
}
