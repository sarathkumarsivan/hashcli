package hash

import (
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
)

func Crc32(text string) []byte {
	tablePolynomial := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(tablePolynomial)
	hash.Write([]byte(text))
	return hash.Sum(nil)[:]
}

// Crc32Hex returns the CRC32 checksum of a text in
// hexadecimal encoding format.
func Crc32Hex(text string) string {
	return hex.EncodeToString(Crc32(text))
}

// Crc32Base64StdEnc returns the MD5 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Crc32Base64StdEnc(text string) string {
	return base64.StdEncoding.EncodeToString(Crc32(text))
}
