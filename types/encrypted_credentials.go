package types

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"

	"github.com/udev-21/golang-tbot-api/types/passport"
	"github.com/udev-21/golang-tbot-api/utils"
)

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

func (ec EncryptedCredentials) Decrypt(privKey *rsa.PrivateKey) (*passport.Credentials, error) {
	credentials_secret, err := utils.RSA_OAEP_Decrypt(ec.Secret, privKey)
	if err != nil {
		return nil, err
	}
	credentials_hash, err := base64.StdEncoding.DecodeString(ec.Hash)
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(ec.Data)
	if err != nil {
		return nil, err
	}
	rawjson, err := utils.DecryptPasswordDataCredentials(credentials_secret, credentials_hash, ciphertext)
	if err != nil {
		return nil, err
	}
	var res passport.Credentials
	if err := json.Unmarshal(rawjson, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
