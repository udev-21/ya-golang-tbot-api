package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
)

// Decrypt - decrypting algorithm according to https://core.telegram.org/passport#datacredentials
// all fields must be bas64 decoded
func DecryptPasswordDataCredentials(secret, hash, encryptedData []byte) ([]byte, error) {
	var data = make([]byte, len(encryptedData))
	copy(data, encryptedData)

	hasher := sha512.New()
	_, err := hasher.Write(append(secret, hash...))
	if err != nil {
		return nil, err
	}
	secret_hash := hasher.Sum(nil)

	key, iv := secret_hash[:32], secret_hash[32:32+16]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("EncryptedData is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	sha256Hasher := sha256.New()
	_, err = sha256Hasher.Write(data)
	if err != nil {
		return nil, err
	}

	checkSum := sha256Hasher.Sum(nil)
	if isNotEqual(checkSum, hash) {
		return nil, errors.New("checksum verification failed")
	}

	return data[data[0]:], nil
}

func isNotEqual(checkSum, hash []byte) bool {
	return base64.StdEncoding.EncodeToString(checkSum) != base64.StdEncoding.EncodeToString(hash)
}
