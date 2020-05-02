package hashcli

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

const (
	Algorithm       = "algorithm"
	AlgorithmMD5    = "md5"
	AlgorithmSHA1   = "sha1"
	AlgorithmSHA256 = "sha256"
	AlgorithmSHA512 = "sha512"
)

func GetHash(text, algorithm string) (string, error) {
	switch algorithm {
	case AlgorithmMD5:
		return getHash(md5.New(), text)
	case AlgorithmSHA1:
		return getHash(sha1.New(), text)
	case AlgorithmSHA256:
		return getHash(sha256.New(), text)
	case AlgorithmSHA512:
		return getHash(sha512.New(), text)
	default:
		return "", nil
	}
}

func getHash(h hash.Hash, text string) (string, error) {
	_, err := h.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func GetFileHash(filename, algorithm string) (string, error) {
	switch algorithm {
	case AlgorithmMD5:
		return getFileHash(md5.New(), filename)
	case AlgorithmSHA1:
		return getFileHash(sha1.New(), filename)
	case AlgorithmSHA256:
		return getFileHash(sha256.New(), filename)
	case AlgorithmSHA512:
		return getFileHash(sha512.New(), filename)
	default:
		return "", nil
	}
}

func getFileHash(h hash.Hash, filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
