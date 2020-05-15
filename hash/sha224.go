package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func SHA224(text string) ([]byte, error) {
	hash := sha256.New224()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
}

func SHA224Hex(text string) (string, error) {
	hash, err := SHA224(text)
	return hex.EncodeToString(hash), err
}

func SHA224Base64StdEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA224Base64URLEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA224Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA224Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA224File(path string) ([]byte, error) {
	hash := sha256.New224()
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

func SHA224FileHex(path string) (string, error) {
	hash, err := SHA224File(path)
	return hex.EncodeToString(hash), err
}

func SHA224FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA224FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA224FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}
