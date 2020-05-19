package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func SHA512(text string) ([]byte, error) {
	hash := sha512.New()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
}

func SHA512Hex(text string) (string, error) {
	hash, err := SHA512(text)
	return hex.EncodeToString(hash), err
}

func SHA512Base64StdEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA512Base64URLEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA512Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA512File(path string) ([]byte, error) {
	hash := sha512.New()
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

func SHA512FileHex(path string) (string, error) {
	hash, err := SHA512File(path)
	return hex.EncodeToString(hash), err
}

func SHA512FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA512File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA512FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA512File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}
