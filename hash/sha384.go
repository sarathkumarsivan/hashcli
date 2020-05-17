package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func SHA384(text string) ([]byte, error) {
	hash := sha512.New384()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
}

func SHA384Hex(text string) (string, error) {
	hash, err := SHA384(text)
	return hex.EncodeToString(hash), err
}

func SHA384Base64StdEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA384Base64URLEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA384Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA384Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA384File(path string) ([]byte, error) {
	hash := sha512.New384()
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

func SHA384FileHex(path string) (string, error) {
	hash, err := SHA384File(path)
	return hex.EncodeToString(hash), err
}

func SHA384FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA384FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA384FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA384FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
