package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func SHA256(text string) ([]byte, error) {
	hash := sha256.New()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
}

func SHA256Hex(text string) (string, error) {
	hash, err := SHA256(text)
	return hex.EncodeToString(hash), err
}

func SHA256Base64StdEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA256Base64URLEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA256File(path string) ([]byte, error) {
	hash := sha256.New()
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func SHA256FileHex(path string) (string, error) {
	hash, err := SHA256File(path)
	return hex.EncodeToString(hash), err
}

func SHA256FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA256FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
