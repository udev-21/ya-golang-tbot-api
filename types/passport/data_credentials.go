package passport

import (
	"encoding/base64"

	"github.com/udev-21/golang-tbot-api/utils"
)

// https://core.telegram.org/passport#datacredentials
type DataCredentials struct {
	DataHash string `json:"data_hash"`
	Secret   string `json:"secret"`
}

// cipherText must be bas64ecoded
func (dc *DataCredentials) Decrypt(cipherText string) ([]byte, error) {
	bsecret, err := base64.StdEncoding.DecodeString(dc.Secret)
	if err != nil {
		return nil, err
	}
	hdataHash, err := base64.StdEncoding.DecodeString(dc.DataHash)
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	return utils.DecryptPasswordDataCredentials(bsecret, hdataHash, data)
}
