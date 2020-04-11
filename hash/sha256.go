package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Sha256(text string) ([]byte, error) {
	hash := sha256.New()
	return hashText(hash, text)
}

func Sha256Hex(text string) (string, error) {
	hash, err := Sha256(text)
	return hex.EncodeToString(hash), err
}

func Sha256Base64StdEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha256Base64URLEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha256Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha256File(path string) ([]byte, error) {
	hash := sha256.New()
	return hashFile(hash, path)
}

func Sha256FileHex(path string) (string, error) {
	hash, err := Sha256File(path)
	return hex.EncodeToString(hash), err
}

func Sha256FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha256FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha256FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha256FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha256Dir(path string) ([]byte, error) {
	hash := sha256.New()
	return hashDir(hash, path)
}

func Sha256DirHex(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return hex.EncodeToString(hash), err
}

func SHA256DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA256DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA256Path(path string) ([]byte, error) {
	hash := sha256.New()
	return hashPath(hash, path)
}

func SHA256PathHex(path string) (string, error) {
	hash, err := SHA256Path(path)
	return hex.EncodeToString(hash), err
}

func SHA256PathBase64StdEnc(path string) (string, error) {
	hash, err := SHA256Path(path)
	return hex.EncodeToString(hash), err
}

func SHA256PathBase64URLEnc(path string) (string, error) {
	hash, err := SHA256Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256PathBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA256Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256PathBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA256Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
