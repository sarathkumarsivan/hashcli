package hash

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HMAC256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}
