package hashcli

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func GetSHA1Hash(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetMD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
