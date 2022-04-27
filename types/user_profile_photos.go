package types

type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}
