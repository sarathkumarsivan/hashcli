package hashcli

import (
	"crypto/sha1"
	"io"
)

func GetSHA1Hash(txt string) string {
	h := sha1.New()
	io.WriteString(h, txt)
	return string(h.Sum(nil))
}
