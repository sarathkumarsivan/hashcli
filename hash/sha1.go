package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
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
