package hash

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

func MD5(text string) ([]byte, error) {
	hash := md5.New()
	_, err := io.WriteString(hash, text)
	return hash.Sum(nil), err
}

func MD5Hex(text string) (string, error) {
	hash, err := MD5(text)
	return hex.EncodeToString(hash), err
}

func MD5Base64StdEnc(text string) (string, error) {
	hash, err := MD5(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func MD5Base64URLEnc(text string) (string, error) {
	hash, err := MD5(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func MD5Base64RawURLEnc(text string) (string, error) {
	hash, err := MD5(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func MD5Base64RawStdEnc(text string) (string, error) {
	hash, err := MD5(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Md5File(path string) ([]byte, error) {
	hash := md5.New()
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
