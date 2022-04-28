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

func (b *BaseInputMedia) Files() []Uploadable {
	if tmp, ok := b.Media.(Uploadable); ok {
		return []Uploadable{tmp}
	}
	return nil
}
func (b *BaseInputMedia) WithMedia(file InputFile) {
	b.Media = file
}

func (b *BaseInputMedia) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(b)
}

const (
	InputMediaPhotoType     = "photo"
	InputMediaVideoType     = "video"
	InputMediaAnimationType = "animation"
	InputMediaAudioType     = "audio"
	InputMediaDocumentType  = "document"
)

func NewInputMediaPhoto(media InputFile) *InputMediaPhoto {
	return &InputMediaPhoto{
		BaseInputMedia: BaseInputMedia{
			Type:  InputMediaPhotoType,
			Media: media,
		},
	}
}

type InputMediaPhoto struct {
	BaseInputMedia
	BaseSendMediaGroupable
}

func NewInputMediaVideo(media InputFile) *InputMediaVideo {
	return &InputMediaVideo{
		BaseInputMedia: BaseInputMedia{
			Type:  InputMediaVideoType,
			Media: media,
		},
	}
}

type InputMediaVideo struct {
	BaseInputMedia
	BaseSendMediaGroupable
	Width             int64 `json:"width,omitempty"`
	Height            int64 `json:"height,omitempty"`
	SupportsStreaming bool  `json:"supports_streaming,omitempty"`

	Thumber
	Durationer
}

func (bmv *InputMediaVideo) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(bmv)
}

func (imv *InputMediaVideo) Files() []Uploadable {
	files := imv.BaseInputMedia.Files()
	if imv.Thumb != nil {
		if tmp, ok := imv.Thumb.(Uploadable); ok {
			files = append(files, tmp)
		}
	}
	return files
}

func NewInputMediaAnimation(media InputFile) *InputMediaAnimation {
	return &InputMediaAnimation{
		BaseInputMedia: BaseInputMedia{
			Type:  InputMediaVideoType,
			Media: media,
		},
	}
}

type InputMediaAnimation struct {
	BaseInputMedia
	Width  *int64 `json:"width,omitempty"`
	Height *int64 `json:"height,omitempty"`
	Durationer
	Thumber
}

func (ima *InputMediaAnimation) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(ima)
}

func (ima *InputMediaAnimation) Files() []Uploadable {
	files := ima.BaseInputMedia.Files()
	if ima.Thumb != nil {
		if tmp, ok := ima.Thumb.(Uploadable); ok {
			files = append(files, tmp)
		}
	}
	return files
}

func NewInputMediaAudio(media InputFile) *InputMediaAudio {
	return &InputMediaAudio{
		BaseInputMedia: BaseInputMedia{
			Type:  InputMediaVideoType,
			Media: media,
		},
	}
}

type InputMediaAudio struct {
	BaseInputMedia
	BaseSendMediaGroupable
	Durationer

	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`
	Thumber
}

func (ima *InputMediaAudio) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(ima)
}

func (ima *InputMediaAudio) Files() []Uploadable {
	files := ima.BaseInputMedia.Files()
	if ima.Thumb != nil {
		if tmp, ok := ima.Thumb.(Uploadable); ok {
			files = append(files, tmp)
		}
	}
	return files
}

func NewInputMediaDocument(media InputFile) *InputMediaDocument {
	return &InputMediaDocument{
		BaseInputMedia: BaseInputMedia{
			Type:  InputMediaDocumentType,
			Media: media,
		},
	}
}

type InputMediaDocument struct {
	BaseInputMedia
	BaseSendMediaGroupable
	Thumber
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

func (document *InputMediaDocument) Params() (Params, error) {
	return utils.ConvertToMapStringInterface(document)
}

func (document *InputMediaDocument) Files() []Uploadable {
	files := document.BaseInputMedia.Files()

	if tmp, ok := document.Thumb.(Uploadable); ok {
		files = append(files, tmp)
	}

	return files
}
