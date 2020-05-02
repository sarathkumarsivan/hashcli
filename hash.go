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
	Algorithm  = "algorithm"
	MD5Hash    = "md5"
	SHA1Hash   = "sha1"
	SHA256Hash = "sha256"
	SHA512Hash = "sha512"
)

func MakeHash(text, algorithm string) (string, error) {
	switch algorithm {
	case MD5Hash:
		return makeTextHash(md5.New(), text)
	case SHA1Hash:
		return makeTextHash(sha1.New(), text)
	case SHA256Hash:
		return makeTextHash(sha256.New(), text)
	case SHA512Hash:
		return makeTextHash(sha512.New(), text)
	default:
		return "", nil
	}
}

func makeTextHash(h hash.Hash, text string) (string, error) {
	_, err := h.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MakeFileHash(path, algorithm string) (string, error) {
	switch algorithm {
	case MD5Hash:
		return makeHash(md5.New(), path)
	case SHA1Hash:
		return makeHash(sha1.New(), path)
	case SHA256Hash:
		return makeHash(sha256.New(), path)
	case SHA512Hash:
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
