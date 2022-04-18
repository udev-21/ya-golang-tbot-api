package types

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}
