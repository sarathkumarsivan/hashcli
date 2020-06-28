package hash

import (
	"encoding/hex"
	"hash/crc32"
	"io"
	"strings"
)

func Crc32cHash(text string) ([]byte, error) {
	tablePolynomial := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, strings.NewReader(text)); err != nil {
		return nil, err
	}
	return hash.Sum(nil)[:], nil
}

func Crc32cHashHex(text string) (string, error) {
	hash, err := Crc32cHash(text)
	return hex.EncodeToString(hash), err
}
