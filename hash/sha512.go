package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func SHA512(text string) ([]byte, error) {
	hash := sha512.New()
	return hashText(hash, text)
}

func SHA512Hex(text string) (string, error) {
	hash, err := SHA512(text)
	return hex.EncodeToString(hash), err
}

func SHA512Base64StdEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA512Base64URLEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA512Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA512(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA512File(path string) ([]byte, error) {
	hash := sha512.New()
	return hashFile(hash, path)
}

func SHA512FileHex(path string) (string, error) {
	hash, err := SHA512File(path)
	return hex.EncodeToString(hash), err
}

func SHA512FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA512File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA512FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA512File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA512File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA512FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA512File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA512Dir(path string) ([]byte, error) {
	hash := sha512.New()
	return hashDir(hash, path)
}

func SHA512DirHex(path string) (string, error) {
	hash, err := SHA512Dir(path)
	return hex.EncodeToString(hash), err
}

func SHA512DirBase64StdEnc(path string) (string, error) {
	hash, err := SHA512Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA512DirBase64URLEnc(path string) (string, error) {
	hash, err := SHA512Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512DirBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA512Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA512DirBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA512Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA512Path(path string) ([]byte, error) {
	hash := sha512.New()
	return hashPath(hash, path)
}

func SHA512PathHex(path string) (string, error) {
	hash, err := SHA512Path(path)
	return hex.EncodeToString(hash), err
}
