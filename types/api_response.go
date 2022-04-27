package types

type ApiResponse struct {
	OK          bool                `json:"ok"`
	Result      interface{}         `json:"result,omitempty"`
	ErrorCode   *int64              `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Description *string             `json:"description,omitempty"`

	RawJson []byte `json:"golangtbotapi_raw,omitempty"`
}
