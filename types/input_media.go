package types

type InputMedia struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// can be either InputFile or string
	// TODO: implement this one
	Thumb           interface{}     `json:"thumb,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Width           *int64          `json:"width,omitempty"`
	Height          *int64          `json:"height,omitempty"`
	Duration        *int64          `json:"duration,omitempty"`

	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`

	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`

	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
}
