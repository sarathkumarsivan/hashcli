package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Hmac256(message string, secret string) string {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
