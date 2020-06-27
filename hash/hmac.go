package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HMAC256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HMAC256Hex(message string, secret string) string {
	bytes := HMAC256(message, secret)
	return hex.EncodeToString(bytes)
}
