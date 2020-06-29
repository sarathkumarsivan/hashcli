package hash

import (
	"encoding/hex"
	"hash/crc32"
)

func Crc32(text string) []byte {
	tablePolynomial := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(tablePolynomial)
	hash.Write([]byte(text))
	return hash.Sum(nil)[:]
}

// Crc32HashHex returns the CRC32 checksum of a text in
// hexadecimal encoding format.
func Crc32HashHex(text string) string {
	hash := Crc32(text)
	return hex.EncodeToString(hash)
}
