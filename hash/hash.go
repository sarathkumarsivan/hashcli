package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
	"path/filepath"
)

var ErrUnsupportedAlgorithm = errors.New("hashutils: unsupported hashing algorithm")

// Algorithm represents the type of hashing algorithms supported by
// this package.
type Algorithm string

const (
	MD5Hash    Algorithm = "md5"
	FNV32Hash            = "fnv32"
	FNV32aHash           = "fnv32a"
	FNV64Hash            = "fnv64"
	FNV64aHash           = "fnv64a"
	SHA1Hash             = "sha1"
	SHA224Hash           = "sha224"
	SHA256Hash           = "sha256"
	SHA512Hash           = "sha512"
)

// An Encoding is a radix 64 encoding/decoding scheme, defined by a
// 64-character alphabet.
type Encoding string

const (
	Hex    Encoding = "hex"
	Base64          = "base64"
	Buffer          = "buffer"
	Binary          = "binary"
)

type ExtHash interface {
	HashText(text string) (string, error)
	HashFile(path string) (string, error)
	HashFiles(paths ...string) (map[string]string, error)
	HashDir(path string) (string, error)
	HashPath(path string) (string, error)
}

type ExtHashBuilder interface {
	Algorithm(Algorithm) ExtHashBuilder
	Encoding(Encoding) ExtHashBuilder
	Build() ExtHash
}

type hashBuilder struct {
	algorithm Algorithm
	encoding  Encoding
}

func (h *hashBuilder) Algorithm(algorithm Algorithm) ExtHashBuilder {
	h.algorithm = algorithm
	return h
}

func (h *hashBuilder) Encoding(encoding Encoding) ExtHashBuilder {
	h.encoding = encoding
	return h
}

func (h *hashBuilder) Build() ExtHash {
	return &hashMaker{
		algorithm: h.algorithm,
		encoding:  h.encoding,
	}
}

func New() ExtHashBuilder {
	return &hashBuilder{}
}

type hashMaker struct {
	algorithm Algorithm
	encoding  Encoding
}

func (m *hashMaker) HashText(text string) (string, error) {
	if m.algorithm == MD5Hash && m.encoding == Hex {
		return hashTextHex(md5.New(), text)
	}
	if m.algorithm == SHA1Hash && m.encoding == Hex {
		return hashTextHex(sha1.New(), text)
	}
	if m.algorithm == SHA256Hash && m.encoding == Hex {
		return hashTextHex(sha256.New(), text)
	}
	if m.algorithm == SHA512Hash && m.encoding == Hex {
		return hashTextHex(sha512.New(), text)
	}
	if m.algorithm == MD5Hash && m.encoding == Base64 {
		return hashTextBase64(md5.New(), text)
	}
	if m.algorithm == SHA1Hash && m.encoding == Base64 {
		return hashTextBase64(sha1.New(), text)
	}
	if m.algorithm == SHA256Hash && m.encoding == Base64 {
		return hashTextBase64(sha256.New(), text)
	}
	if m.algorithm == SHA512Hash && m.encoding == Base64 {
		return hashTextBase64(sha512.New(), text)
	}
	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashFile(path string) (string, error) {
	if m.algorithm == MD5Hash && m.encoding == Hex {
		return hashFileHex(md5.New(), path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Hex {
		return hashFileHex(sha1.New(), path)
	}
	if m.algorithm == SHA256Hash && m.encoding == Hex {
		return hashFileHex(sha256.New(), path)
	}
	if m.algorithm == SHA512Hash && m.encoding == Hex {
		return hashFileHex(sha512.New(), path)
	}
	if m.algorithm == MD5Hash && m.encoding == Base64 {
		return hashFileBase64(md5.New(), path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Base64 {
		return hashFileBase64(sha1.New(), path)
	}
	if m.algorithm == SHA256Hash && m.encoding == Base64 {
		return hashFileBase64(sha256.New(), path)
	}
	if m.algorithm == SHA512Hash && m.encoding == Base64 {
		return hashFileBase64(sha512.New(), path)
	}
	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashFiles(paths ...string) (map[string]string, error) {
	pathHashes := make(map[string]string, len(paths))
	for _, path := range paths {
		if m.algorithm == MD5Hash && m.encoding == Hex {
			hash, err := hashFileHex(md5.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA1Hash && m.encoding == Hex {
			hash, err := hashFileHex(sha1.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA256Hash && m.encoding == Hex {
			hash, err := hashFileHex(sha256.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA512Hash && m.encoding == Hex {
			hash, err := hashFileHex(sha512.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == MD5Hash && m.encoding == Base64 {
			hash, err := hashFileBase64(md5.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA1Hash && m.encoding == Base64 {
			hash, err := hashFileBase64(sha1.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA256Hash && m.encoding == Base64 {
			hash, err := hashFileBase64(sha256.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA512Hash && m.encoding == Base64 {
			hash, err := hashFileBase64(sha512.New(), path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
	}
	return pathHashes, nil
}

func (m *hashMaker) HashDir(path string) (string, error) {
	if m.algorithm == MD5Hash && m.encoding == Hex {
		return hashDirHex(md5.New(), path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Hex {
		return hashDirHex(sha1.New(), path)
	}
	if m.algorithm == SHA256Hash && m.encoding == Hex {
		return hashDirHex(sha256.New(), path)
	}
	if m.algorithm == SHA512Hash && m.encoding == Hex {
		return hashDirHex(sha512.New(), path)
	}
	if m.algorithm == MD5Hash && m.encoding == Base64 {
		return hashDirBase64(md5.New(), path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Base64 {
		return hashDirBase64(sha1.New(), path)
	}
	if m.algorithm == SHA256Hash && m.encoding == Base64 {
		return hashDirBase64(sha256.New(), path)
	}
	if m.algorithm == SHA512Hash && m.encoding == Base64 {
		return hashDirBase64(sha512.New(), path)
	}
	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashPath(path string) (string, error) {
	return "", ErrUnsupportedAlgorithm
}

func hashTextHex(hash hash.Hash, text string) (string, error) {
	bytes, err := hashText(hash, text)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hashTextBase64(hash hash.Hash, text string) (string, error) {
	bytes, err := hashText(hash, text)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
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

func hashFileBase64(hash hash.Hash, text string) (string, error) {
	bytes, err := hashFile(hash, text)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
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

func hashPath(hash hash.Hash, path string) ([]byte, error) {
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

func hashDirBase64(hash hash.Hash, text string) (string, error) {
	bytes, err := hashDir(hash, text)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
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
