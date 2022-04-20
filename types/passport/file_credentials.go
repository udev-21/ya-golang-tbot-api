package passport

import (
	"crypto/rsa"
)

type FileCredentials struct {
	FileHash string `json:"file_hash"`
	Secret   string `json:"secret"`
}

func (dc *FileCredentials) Decrypt(cipherText string, privKey *rsa.PrivateKey) ([]byte, error) {

	return nil, nil
}
