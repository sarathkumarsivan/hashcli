package hash

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func SHA1(text string) ([]byte, error) {
	hash := sha1.New()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
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

func SHA1FileHex(path string) (string, error) {
	hash, err := SHA1File(path)
	return hex.EncodeToString(hash), err
}
