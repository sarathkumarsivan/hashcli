package hashcli

import (
	"crypto/sha1"
	"io"
)

func GetSHA1Hash(text string) string {
	h := sha1.New()
	io.WriteString(h, text)
	return string(h.Sum(nil))
}
