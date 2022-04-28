package types

import "encoding/json"

type ApiResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	ErrorCode   *int64              `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Description *string             `json:"description,omitempty"`
}
