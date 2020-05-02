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

func GetHash(text, hash string) (string, error) {
	switch hash {
	case MD5Hash:
		return MakeHash(md5.New(), text)
	case SHA1Hash:
		return MakeHash(sha1.New(), text)
	case SHA256Hash:
		return MakeHash(sha256.New(), text)
	case SHA512Hash:
		return MakeHash(sha512.New(), text)
	default:
		return "", nil
	}
}

func MakeHash(hash hash.Hash, text string) (string, error) {
	_, err := hash.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GetFileHash(path, hash string) (string, error) {
	switch hash {
	case MD5Hash:
		return MakeFileHash(md5.New(), path)
	case SHA1Hash:
		return MakeFileHash(sha1.New(), path)
	case SHA256Hash:
		return MakeFileHash(sha256.New(), path)
	case SHA512Hash:
		return MakeFileHash(sha512.New(), path)
	default:
		return "", nil
	}
}

func MakeFileHash(hash hash.Hash, path string) (string, error) {
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

func makeDirHash(hash hash.Hash, path string) (string, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		io.WriteString(hash, path)
		return nil
	})

	if err != nil {
		return "", nil
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
