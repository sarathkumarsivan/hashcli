package hash

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func md5Hash(text string) ([]byte, error) {
	hash := md5.New()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
}

func md5Hex(text string) (string, error) {
	hash, err := md5Hash(text)
	return hex.EncodeToString(hash), err
}

func md5Base64StdEnc(text string) (string, error) {
	hash, err := md5Hash(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func md5Base64URLEnc(text string) (string, error) {
	hash, err := md5Hash(text)
	return base64.URLEncoding.EncodeToString(hash), err
}
func md5Base64RawURLEnc(text string) (string, error) {
	hash, err := md5Hash(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func md5Base64RawStdEnc(text string) (string, error) {
	hash, err := md5Hash(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
