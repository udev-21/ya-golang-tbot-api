package types

type ApiResponse struct {
	OK          bool                `json:"ok"`
	Result      interface{}         `json:"result,omitempty"`
	ErrorCode   *int64              `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Description *string             `json:"description,omitempty"`
	// result not always decoded as we want i.e. large int may encoded as float,
	// so this field reversed for decoding specific structs
	RawJson []byte `json:"golangtbotapi_raw,omitempty"`
}
