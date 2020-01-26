package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// Sha512 returns SHA-512 checksum of a text as bytes.
func Sha512(text string) ([]byte, error) {
	hash := sha512.New()
	return hashText(hash, text)
}

// Sha512Hex returns the SHA-512 checksum of a text in
// hexadecimal encoding format.
func Sha512Hex(text string) (string, error) {
	hash, err := Sha512(text)
	return hex.EncodeToString(hash), err
}

// Sha512Base64StdEnc returns the SHA-512 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Sha512Base64StdEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha512Base64URLEnc returns the SHA-512 checksum of a text in
// an alternate base64 encoding defined in RFC 4648.
func Sha512Base64URLEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

// SHA512Base64RawURLEnc returns the SHA-512 checksum of a text in
// a padded alternate base64 encoding defined in RFC 4648.
func SHA512Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Sha512Base64RawStdEnc returns the SHA-512 checksum of a text in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Sha512Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha512(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Sha512File returns SHA-512 checksum of a file as bytes.
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

func Sha512DirHex(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return hex.EncodeToString(hash), err
}

func Sha512DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func Sha512DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha512DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha512DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha512Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha512Path(path string) ([]byte, error) {
	hash := sha512.New()
	return hashPath(hash, path)
}

func Sha512PathHex(path string) (string, error) {
	hash, err := Sha512Path(path)
	return hex.EncodeToString(hash), err
}

func Sha512PathBase64StdEnc(path string) (string, error) {
	hash, err := Sha512Path(path)
	return hex.EncodeToString(hash), err
}

func Sha512PathBase64URLEnc(path string) (string, error) {
	hash, err := Sha512Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha512PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha512Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha512PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha512Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
