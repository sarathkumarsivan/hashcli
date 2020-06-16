package hash

import (
	"errors"
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
	SHA256Hash           = "sha256"
	SHA224Hash           = "sha224"
	SHA512Hash           = "sha512"
	SHA384Hash           = "sha384"
)

// An Encoding is a radix 64 encoding/decoding scheme, defined by a
// 64-character alphabet.
type Encoding string

const (
	Hex    Encoding = "hex"
	Base64          = "base64"
)

type ExtHash interface {
	HashText(text string) (string, error)
	HashFile(path string) (string, error)
	HashFiles(paths ...string) (map[string]string, error)
	// HashDir(path string) (string, error)
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
		return Md5Hex(text)
	}
	if m.algorithm == SHA1Hash && m.encoding == Hex {
		return SHA1Hex(text)
	}
	if m.algorithm == SHA224Hash && m.encoding == Hex {
		return SHA224Hex(text)
	}
	if m.algorithm == SHA256Hash && m.encoding == Hex {
		return SHA256Hex(text)
	}
	if m.algorithm == SHA384Hash && m.encoding == Hex {
		return SHA384Hex(text)
	}
	if m.algorithm == SHA512Hash && m.encoding == Hex {
		return SHA512Hex(text)
	}

	if m.algorithm == MD5Hash && m.encoding == Base64 {
		return MD5Base64StdEnc(text)
	}
	if m.algorithm == SHA1Hash && m.encoding == Base64 {
		return SHA1Base64StdEnc(text)
	}
	if m.algorithm == SHA224Hash && m.encoding == Base64 {
		return SHA224Base64StdEnc(text)
	}
	if m.algorithm == SHA256Hash && m.encoding == Base64 {
		return SHA256Base64StdEnc(text)
	}
	if m.algorithm == SHA384Hash && m.encoding == Base64 {
		return SHA384Base64StdEnc(text)
	}
	if m.algorithm == SHA512Hash && m.encoding == Base64 {
		return SHA512Base64StdEnc(text)
	}

	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashFile(path string) (string, error) {
	if m.algorithm == MD5Hash && m.encoding == Hex {
		return MD5FileHex(path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Hex {
		return SHA1FileHex(path)
	}
	if m.algorithm == SHA224Hash && m.encoding == Hex {
		return SHA224FileHex(path)
	}
	if m.algorithm == SHA256Hash && m.encoding == Hex {
		return SHA256FileHex(path)
	}
	if m.algorithm == SHA384Hash && m.encoding == Hex {
		return SHA384FileHex(path)
	}
	if m.algorithm == SHA512Hash && m.encoding == Hex {
		return SHA512FileHex(path)
	}

	if m.algorithm == MD5Hash && m.encoding == Base64 {
		return MD5FileBase64StdEnc(path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Base64 {
		return SHA1FileBase64StdEnc(path)
	}
	if m.algorithm == SHA224Hash && m.encoding == Base64 {
		return SHA224FileBase64StdEnc(path)
	}
	if m.algorithm == SHA256Hash && m.encoding == Base64 {
		return SHA256FileBase64StdEnc(path)
	}
	if m.algorithm == SHA384Hash && m.encoding == Base64 {
		return SHA384FileBase64StdEnc(path)
	}
	if m.algorithm == SHA512Hash && m.encoding == Base64 {
		return SHA512FileBase64StdEnc(path)
	}

	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashFiles(paths ...string) (map[string]string, error) {
	pathHashes := make(map[string]string, len(paths))
	for _, path := range paths {
		if m.algorithm == MD5Hash && m.encoding == Hex {
			hash, err := MD5FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA1Hash && m.encoding == Hex {
			hash, err := SHA1FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA224Hash && m.encoding == Hex {
			hash, err := SHA224FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA256Hash && m.encoding == Hex {
			hash, err := SHA256FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA384Hash && m.encoding == Hex {
			hash, err := SHA384FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA512Hash && m.encoding == Hex {
			hash, err := SHA512FileHex(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}

		if m.algorithm == MD5Hash && m.encoding == Base64 {
			hash, err := MD5FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA1Hash && m.encoding == Base64 {
			hash, err := SHA1FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA224Hash && m.encoding == Base64 {
			hash, err := SHA224FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA256Hash && m.encoding == Base64 {
			hash, err := SHA256FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA384Hash && m.encoding == Base64 {
			hash, err := SHA384FileBase64StdEnc(path)
			if err != nil {
				return pathHashes, err
			}
			pathHashes[path] = hash
		}
		if m.algorithm == SHA512Hash && m.encoding == Base64 {
			hash, err := SHA512FileBase64StdEnc(path)
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
		return MD5DirHex(path)
	}
	if m.algorithm == SHA1Hash && m.encoding == Hex {
		return SHA1DirHex(path)
	}
	return "", ErrUnsupportedAlgorithm
}

func (m *hashMaker) HashPath(path string) (string, error) {
	return "", ErrUnsupportedAlgorithm
}
