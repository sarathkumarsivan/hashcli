package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SHA224(text string) ([]byte, error) {
	hash := sha256.New224()
	return hashText(hash, text)
}

func SHA224Hex(text string) (string, error) {
	hash, err := SHA224(text)
	return hex.EncodeToString(hash), err
}

func SHA224Base64StdEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA224Base64URLEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA224Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA224Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA224(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA224File(path string) ([]byte, error) {
	hash := sha256.New224()
	return hashFile(hash, path)
}

func SHA224FileHex(path string) (string, error) {
	hash, err := SHA224File(path)
	return hex.EncodeToString(hash), err
}

func SHA224FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA224FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA224FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA224FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA224File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA224Dir(path string) ([]byte, error) {
	hash := sha256.New224()
	return hashDir(hash, path)
}

func SHA224DirHex(path string) (string, error) {
	hash, err := SHA224Dir(path)
	return hex.EncodeToString(hash), err
}

func SHA224DirBase64StdEnc(path string) (string, error) {
	hash, err := SHA224Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA224DirBase64URLEnc(path string) (string, error) {
	hash, err := SHA224Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA224DirBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA224Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA224DirBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA224Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA224Path(path string) ([]byte, error) {
	hash := sha256.New224()
	return hashPath(hash, path)
}

func SHA224PathHex(path string) (string, error) {
	hash, err := SHA224Path(path)
	return hex.EncodeToString(hash), err
}

func SHA224PathBase64StdEnc(path string) (string, error) {
	hash, err := SHA224Path(path)
	return hex.EncodeToString(hash), err
}
