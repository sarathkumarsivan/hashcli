package hashcli

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
)

var ErrUnsupportedHash = errors.New("unsupported hashing algorithm")

type Algorithm string

const (
	MD5Hash    Algorithm = "md5"
	SHA1Hash             = "sha1"
	SHA256Hash           = "sha256"
	SHA512Hash           = "sha512"
)

type Encoding string

const (
	Hex    Encoding = "hex"
	Base64          = "base64"
	Buffer          = "buffer"
	Binary          = "binary"
)

type HashMaker interface {
	HashText(text string) (string, error)
	HashFile(path string) (string, error)
	HashFiles(path ...string) (string, error)
	HashDir(path string) (string, error)
}

type HashBuilder interface {
	Algorithm(Algorithm) HashBuilder
	Encoding(Encoding) HashBuilder
	Build() HashMaker
}

type hashBuilder struct {
	algorithm Algorithm
	encoding  Encoding
}

func (builder *hashBuilder) Algorithm(algorithm Algorithm) HashBuilder {
	builder.algorithm = algorithm
	return builder
}

func (builder *hashBuilder) Encoding(encoding Encoding) HashBuilder {
	builder.encoding = encoding
	return builder
}

func (builder *hashBuilder) Build() HashMaker {
	return &hashMaker{
		algorithm: builder.algorithm,
		encoding:  builder.encoding,
	}
}

func New() HashBuilder {
	return &hashBuilder{}
}

type hashMaker struct {
	algorithm Algorithm
	encoding  Encoding
}

func (maker *hashMaker) HashText(text string) (string, error) {
	if maker.algorithm == MD5Hash && maker.encoding == Hex {
		return hashTextHex(md5.New(), text)
	}
	if maker.algorithm == SHA1Hash && maker.encoding == Hex {
		return hashTextHex(sha1.New(), text)
	}
	if maker.algorithm == SHA256Hash && maker.encoding == Hex {
		return hashTextHex(sha256.New(), text)
	}
	if maker.algorithm == SHA512Hash && maker.encoding == Hex {
		return hashTextHex(sha512.New(), text)
	}
	return "", ErrUnsupportedHash
}

func (maker *hashMaker) HashFile(path string) (string, error) {
	if maker.algorithm == MD5Hash && maker.encoding == Hex {
		return hashFileHex(md5.New(), path)
	}
	if maker.algorithm == SHA1Hash && maker.encoding == Hex {
		return hashFileHex(sha1.New(), path)
	}
	if maker.algorithm == SHA256Hash && maker.encoding == Hex {
		return hashFileHex(sha256.New(), path)
	}
	if maker.algorithm == SHA512Hash && maker.encoding == Hex {
		return hashFileHex(sha512.New(), path)
	}
	return "", ErrUnsupportedHash
}

func (maker *hashMaker) HashFiles(path ...string) (string, error) {
	return "TODO:implement", nil
}

func (maker *hashMaker) HashDir(path string) (string, error) {
	if maker.algorithm == MD5Hash && maker.encoding == Hex {
		return hashDirHex(md5.New(), path)
	}
	if maker.algorithm == SHA1Hash && maker.encoding == Hex {
		return hashDirHex(sha1.New(), path)
	}
	if maker.algorithm == SHA256Hash && maker.encoding == Hex {
		return hashDirHex(sha256.New(), path)
	}
	if maker.algorithm == SHA512Hash && maker.encoding == Hex {
		return hashDirHex(sha512.New(), path)
	}
	return "", ErrUnsupportedHash
}

func hashTextHex(hash hash.Hash, text string) (string, error) {
	bytes, err := hashText(hash, text)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hashText(hash hash.Hash, text string) ([]byte, error) {
	_, err := hash.Write([]byte(text))
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func hashFileHex(hash hash.Hash, text string) (string, error) {
	bytes, err := hashFile(hash, text)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hashFile(hash hash.Hash, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func hashFileOrDir(hash hash.Hash, path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	switch mode := info.Mode(); {
	case mode.IsDir():
		return hashFile(hash, path)
	case mode.IsRegular():
		return hashDir(hash, path)
	}
	return nil, nil // TODO: Handle Error
}

func hashDirHex(hash hash.Hash, text string) (string, error) {
	bytes, err := hashDir(hash, text)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hashDir(hash hash.Hash, path string) ([]byte, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		io.WriteString(hash, path)
		return nil
	})
	if err != nil {
		return nil, nil
	}
	return hash.Sum(nil), nil
}

func main() {
	hashcli := New().Algorithm(MD5Hash).Encoding(Hex).Build()
	fmt.Println(hashcli.HashText(""))
	fmt.Println(hashcli.HashFile(""))
}
