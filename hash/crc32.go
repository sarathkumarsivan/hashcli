package hash

import (
	"encoding/hex"
	"hash/crc32"
	"io"
	"strings"
)

func Crc32(text string) ([]byte, error) {
	tablePolynomial := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, strings.NewReader(text)); err != nil {
		return nil, err
	}
	return hash.Sum(nil)[:], nil
}

// Crc32HashHex returns the CRC32 checksum of a text in
// hexadecimal encoding format.
func Crc32HashHex(text string) (string, error) {
	hash, err := Crc32(text)
	return hex.EncodeToString(hash), err
}
