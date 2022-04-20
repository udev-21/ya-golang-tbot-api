package types

type ApiResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}
