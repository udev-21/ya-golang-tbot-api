package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"sort"
)

// check algorith according to: https://core.telegram.org/widgets/login#checking-authorization
func CheckTelegramHash(botToken string, fields map[string]string) bool {
	strs := []string{}
	var hash = ""
	for k, v := range fields {
		if k == "hash" {
			hash = v
			continue
		}
		strs = append(strs, k+"="+v)
	}
	sort.Strings(strs)
	var imploded = ""
	for _, s := range strs {
		if imploded != "" {
			imploded += "\n"
		}
		imploded += s
	}
	sha256hash := sha256.New()
	io.WriteString(sha256hash, botToken)
	hmachash := hmac.New(sha256.New, sha256hash.Sum(nil))
	io.WriteString(hmachash, imploded)
	ss := hex.EncodeToString(hmachash.Sum(nil))
	return hash == ss
}

// encryptedText must be base64 encoded
func RSA_OAEP_Decrypt(encryptedText string, privKey *rsa.PrivateKey) ([]byte, error) {
	if privKey == nil {
		return nil, errors.New("private key nil")
	}
	ct, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return nil, err
	}
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, privKey, ct, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
