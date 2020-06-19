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

func HMAC256Hex(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return hex.EncodeToString(bytes)
}

func HMAC256Base64StdEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HMAC256Base64RawStdEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HMAC256Base64URLEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HMAC256Base64RawURLEnc(message string, secret string) string {
	bytes := Hmac256(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HMAC224(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New224, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HMAC224Hex(message string, secret string) string {
	bytes := HMAC224(message, secret)
	return hex.EncodeToString(bytes)
}

func HMAC224Base64StdEnc(message string, secret string) string {
	bytes := HMAC224(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HMAC224Base64RawStdEnc(message string, secret string) string {
	bytes := HMAC224(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HMAC224Base64URLEnc(message string, secret string) string {
	bytes := HMAC224(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HMAC224Base64RawURLEnc(message string, secret string) string {
	bytes := HMAC224(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HMAC512(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HMAC512Hex(message string, secret string) string {
	bytes := HMAC512(message, secret)
	return hex.EncodeToString(bytes)
}

func HMAC512Base64StdEnc(message string, secret string) string {
	bytes := HMAC512(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HMAC512Base64RawStdEnc(message string, secret string) string {
	bytes := HMAC512(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HMAC512Base64URLEnc(message string, secret string) string {
	bytes := HMAC512(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HMAC512Base64RawURLEnc(message string, secret string) string {
	bytes := HMAC512(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func HMAC384(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New384, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HMAC384Hex(message string, secret string) string {
	bytes := HMAC384(message, secret)
	return hex.EncodeToString(bytes)
}

func HMAC384Base64StdEnc(message string, secret string) string {
	bytes := HMAC384(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HMAC384Base64RawStdEnc(message string, secret string) string {
	bytes := HMAC384(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HMAC384Base64URLEnc(message string, secret string) string {
	bytes := HMAC384(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HMAC384Base64RawURLEnc(message string, secret string) string {
	bytes := HMAC384(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}
