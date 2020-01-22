package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// Sha384 returns SHA-384 checksum of a text as bytes.
func Sha384(text string) ([]byte, error) {
	hash := sha512.New384()
	return hashText(hash, text)
}

// Sha384Hex returns the SHA-384 checksum of a text in
// hexadecimal encoding format.
func Sha384Hex(text string) (string, error) {
	hash, err := Sha384(text)
	return hex.EncodeToString(hash), err
}

func Sha384Base64StdEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha384Base64URLEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha384Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha384Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha384(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha384File(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashFile(hash, path)
}

func Sha384FileHex(path string) (string, error) {
	hash, err := Sha384File(path)
	return hex.EncodeToString(hash), err
}

func Sha384FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha384FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha384FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha384FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha384File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha384Dir(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashDir(hash, path)
}

func Sha384DirHex(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return hex.EncodeToString(hash), err
}

func Sha384DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha384DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha384DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA384DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha384Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha384Path(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashPath(hash, path)
}

func Sha384PathHex(path string) (string, error) {
	hash, err := Sha384Path(path)
	return hex.EncodeToString(hash), err
}

func Sha384PathBase64StdEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return hex.EncodeToString(hash), err
}

func Sha384PathBase64URLEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha384PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha384PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha384Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
