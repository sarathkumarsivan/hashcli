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

func HMAC256Base64StdEnc(message string, secret string) string {
	bytes := HMAC256(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HMAC256Base64URLEnc(message string, secret string) string {
	bytes := HMAC256(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HMAC256Base64RawURLEnc(message string, secret string) string {
	bytes := HMAC256(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HMAC256Base64RawStdEnc(message string, secret string) string {
	bytes := HMAC256(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HMAC224(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New224, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}
