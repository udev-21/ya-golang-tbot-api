package types

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}
