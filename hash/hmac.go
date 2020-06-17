package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func HmacMd5(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(md5.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacMd5Hex(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacMd5Base64StdEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacMd5Base64RawStdEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacMd5Base64URLEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacMd5Base64RawURLEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func Hmac1(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func Hmac1Hex(message string, secret string) string {
	bytes := Hmac1(message, secret)
	return hex.EncodeToString(bytes)
}

func Hmac1Base64StdEnc(message string, secret string) string {
	bytes := Hmac1(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Hmac1Base64RawStdEnc(message string, secret string) string {
	bytes := Hmac1(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func Hmac1Base64URLEnc(message string, secret string) string {
	bytes := Hmac1(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func Hmac1Base64RawURLEnc(message string, secret string) string {
	bytes := Hmac1(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func Hmac256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func Hmac256Hex(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return hex.EncodeToString(bytes)
}

func Hmac256Base64StdEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Hmac256Base64RawStdEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func Hmac256Base64URLEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func Hmac256Base64RawURLEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func Hmac224(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New224, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func Hmac224Hex(message string, secret string) string {
	bytes := Hmac224(message, secret)
	return hex.EncodeToString(bytes)
}

func Hmac224Base64StdEnc(message string, secret string) string {
	bytes := Hmac224(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Hmac224Base64RawStdEnc(message string, secret string) string {
	bytes := Hmac224(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func Hmac224Base64URLEnc(message string, secret string) string {
	bytes := Hmac224(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func Hmac224Base64RawURLEnc(message string, secret string) string {
	bytes := Hmac224(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func Hmac512(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func Hmac512Hex(message string, secret string) string {
	bytes := Hmac512(message, secret)
	return hex.EncodeToString(bytes)
}

func Hmac512Base64StdEnc(message string, secret string) string {
	bytes := Hmac512(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Hmac512Base64RawStdEnc(message string, secret string) string {
	bytes := Hmac512(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func Hmac512Base64URLEnc(message string, secret string) string {
	bytes := Hmac512(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func Hmac512Base64RawURLEnc(message string, secret string) string {
	bytes := Hmac512(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func Hmac384(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New384, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func Hmac384Hex(message string, secret string) string {
	bytes := Hmac384(message, secret)
	return hex.EncodeToString(bytes)
}

func Hmac384Base64StdEnc(message string, secret string) string {
	bytes := Hmac384(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Hmac384Base64RawStdEnc(message string, secret string) string {
	bytes := Hmac384(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func Hmac384Base64URLEnc(message string, secret string) string {
	bytes := Hmac384(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func Hmac384Base64RawURLEnc(message string, secret string) string {
	bytes := Hmac384(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}
