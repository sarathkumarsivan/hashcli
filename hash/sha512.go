package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func Sha512(text string) ([]byte, error) {
	hash := sha512.New()
	return hashText(hash, text)
}

func Sha512Hex(text string) (string, error) {
	hash, err := Sha512(text)
	return hex.EncodeToString(hash), err
}

func Sha512Base64StdEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha512Base64URLEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha512Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha512File(path string) ([]byte, error) {
	hash := sha512.New()
	return hashFile(hash, path)
}

func Sha512FileHex(path string) (string, error) {
	hash, err := Sha512File(path)
	return hex.EncodeToString(hash), err
}

func Sha512FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha512File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha512FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha512File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha512FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha512File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha512FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha512File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha512Dir(path string) ([]byte, error) {
	hash := sha512.New()
	return hashDir(hash, path)
}

func SHA512DirHex(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return hex.EncodeToString(hash), err
}

func SHA512DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA512DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA512DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
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

func SHA512PathBase64StdEnc(path string) (string, error) {
	hash, err := SHA512Path(path)
	return hex.EncodeToString(hash), err
}

func SHA512PathBase64URLEnc(path string) (string, error) {
	hash, err := SHA512Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA512PathBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA512Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA512PathBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA512Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
