package hash

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

// Sha1 returns SHA-1 checksum of a text as bytes.
func Sha1(text string) ([]byte, error) {
	hash := sha1.New()
	return hashText(hash, text)
}

// Sha1Hex returns the SHA-1 checksum of a text in
// hexadecimal encoding format.
func Sha1Hex(text string) (string, error) {
	hash, err := Sha1(text)
	return hex.EncodeToString(hash), err
}

// Sha1Base64StdEnc returns the SHA-1 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Sha1Base64StdEnc(text string) (string, error) {
	hash, err := Sha1(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha1Base64URLEnc returns the SHA-1 checksum of a text in
// an alternate base64 encoding defined in RFC 4648.
func Sha1Base64URLEnc(text string) (string, error) {
	hash, err := Sha1(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha1Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha1(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha1Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha1(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha1File(path string) ([]byte, error) {
	hash := sha1.New()
	return hashFile(hash, path)
}

func Sha1FileHex(path string) (string, error) {
	hash, err := Sha1File(path)
	return hex.EncodeToString(hash), err
}

func Sha1FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha1File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha1FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha1File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha1FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha1File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha1FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha1File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha1Dir(path string) ([]byte, error) {
	hash := sha1.New()
	return hashDir(hash, path)
}

func Sha1DirHex(path string) (string, error) {
	hash, err := Sha1Dir(path)
	return hex.EncodeToString(hash), err
}

func Sha1DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha1Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha1DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha1Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha1DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha1Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha1DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha1Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha1Path(path string) ([]byte, error) {
	hash := sha1.New()
	return hashPath(hash, path)
}

func Sha1PathHex(path string) (string, error) {
	hash, err := Sha1Path(path)
	return hex.EncodeToString(hash), err
}

func Sha1PathBase64StdEnc(path string) (string, error) {
	hash, err := Sha1Path(path)
	return hex.EncodeToString(hash), err
}

func Sha1PathBase64URLEnc(path string) (string, error) {
	hash, err := Sha1Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha1PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha1Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha1PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha1Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
