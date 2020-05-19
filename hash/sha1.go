package hash

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

func SHA1(text string) ([]byte, error) {
	hash := sha1.New()
	return hashText(hash, text)
}

func SHA1Hex(text string) (string, error) {
	hash, err := SHA1(text)
	return hex.EncodeToString(hash), err
}

func SHA1Base64StdEnc(text string) (string, error) {
	hash, err := SHA1(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA1Base64URLEnc(text string) (string, error) {
	hash, err := SHA1(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA1Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA1(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA1Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA1(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA1File(path string) ([]byte, error) {
	hash := sha1.New()
	return hashFile(hash, path)
}

func SHA1FileHex(path string) (string, error) {
	hash, err := SHA1File(path)
	return hex.EncodeToString(hash), err
}

func SHA1FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA1File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA1FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA1File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA1FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA1File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA1FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA1File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
