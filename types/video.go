package types

type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}
