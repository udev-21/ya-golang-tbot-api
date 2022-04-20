package passport

type Credentials struct {
	SecureData map[string]SecureValue `json:"secure_data"`
	Nonce      string                 `json:"nonce"`
}
