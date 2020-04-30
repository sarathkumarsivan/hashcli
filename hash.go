package hashcli

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func GetSHA1Hash(txt string) string {
	h := sha1.New()
	io.WriteString(h, txt)
	return string(h.Sum(nil))
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
