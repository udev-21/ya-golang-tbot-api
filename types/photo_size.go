package types

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     *int64 `json:"file_size,omitempty"`
}
