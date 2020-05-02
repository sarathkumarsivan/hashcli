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
	"path/filepath"
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

func GetFileHash(path, algorithm string) (string, error) {
	switch algorithm {
	case AlgorithmMD5:
		return makeHash(md5.New(), path)
	case AlgorithmSHA1:
		return makeHash(sha1.New(), path)
	case AlgorithmSHA256:
		return makeHash(sha256.New(), path)
	case AlgorithmSHA512:
		return makeHash(sha512.New(), path)
	default:
		return "", nil
	}
}

func makeHash(hash hash.Hash, path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	switch mode := info.Mode(); {
	case mode.IsDir():
		return makeDirHash(hash, path)
	case mode.IsRegular():
		return makeFileHash(hash, path)
	}
	return "", nil
}

func makeFileHash(hash hash.Hash, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func makeDirHash(h hash.Hash, path string) (string, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		io.WriteString(h, path)
		return nil
	})

	if err != nil {
		return "", nil
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
