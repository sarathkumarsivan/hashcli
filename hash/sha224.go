package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
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
