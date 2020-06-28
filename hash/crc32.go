package hash

import (
	"hash/crc32"
	"io"
	"strings"
)

func Crc32cHash(text string, polynomial uint32) ([]byte, error) {
	tablePolynomial := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, strings.NewReader(text)); err != nil {
		return nil, err
	}
	return hash.Sum(nil)[:], nil
}
