package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Sha256 returns SHA-256 checksum of a text as bytes.
func Sha256(text string) ([]byte, error) {
	hash := sha256.New()
	return hashText(hash, text)
}

// Sha256Hex returns the SHA-256 checksum of a text in
// hexadecimal encoding format.
func Sha256Hex(text string) (string, error) {
	hash, err := Sha256(text)
	return hex.EncodeToString(hash), err
}

// Sha256Base64StdEnc returns the SHA-256 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Sha256Base64StdEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha256Base64URLEnc returns the SHA-256 checksum of a text in
// an alternate base64 encoding defined in RFC 4648.
func Sha256Base64URLEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

// SHA256Base64RawURLEnc returns the SHA-256 checksum of a text in
// a padded alternate base64 encoding defined in RFC 4648.
func SHA256Base64RawURLEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Sha256Base64RawStdEnc returns the SHA-256 checksum of a text in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Sha256Base64RawStdEnc(text string) (string, error) {
	hash, err := Sha256(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Sha256File returns SHA-256 checksum of a file as bytes.
func Sha256File(path string) ([]byte, error) {
	hash := sha256.New()
	return hashFile(hash, path)
}

// Sha256FileHex returns the SHA-256 checksum of a file in
// hexadecimal encoding format.
func Sha256FileHex(path string) (string, error) {
	hash, err := Sha256File(path)
	return hex.EncodeToString(hash), err
}

// Sha256FileBase64StdEnc returns the SHA-256 checksum of a file in
// standard base64 encoding, as defined in RFC 4648.
func Sha256FileBase64StdEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha256FileBase64URLEnc returns the SHA-256 checksum of a file in
// an alternate base64 encoding defined in RFC 4648.
func Sha256FileBase64URLEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Sha256FileBase64RawURLEnc returns the SHA-256 checksum of a file in
// a padded alternate base64 encoding defined in RFC 4648.
func Sha256FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Sha256FileBase64RawStdEnc returns the SHA-256 checksum of a file in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Sha256FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha256File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Sha256Dir returns SHA-256 checksum of a directory as bytes.
func Sha256Dir(path string) ([]byte, error) {
	hash := sha256.New()
	return hashDir(hash, path)
}

// Sha256DirHex returns the SHA-256 checksum of a directory in
// hexadecimal encoding format.
func Sha256DirHex(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return hex.EncodeToString(hash), err
}

// Sha256DirBase64StdEnc returns the SHA-256 checksum of a directory in
// standard base64 encoding, as defined in RFC 4648.
func Sha256DirBase64StdEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Sha256DirBase64URLEnc returns the SHA-256 checksum of a directory in
// an alternate base64 encoding defined in RFC 4648.
func Sha256DirBase64URLEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha256DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha256DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha256Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func Sha256Path(path string) ([]byte, error) {
	hash := sha256.New()
	return hashPath(hash, path)
}

func Sha256PathHex(path string) (string, error) {
	hash, err := Sha256Path(path)
	return hex.EncodeToString(hash), err
}

func Sha256PathBase64StdEnc(path string) (string, error) {
	hash, err := Sha256Path(path)
	return hex.EncodeToString(hash), err
}

func Sha256PathBase64URLEnc(path string) (string, error) {
	hash, err := Sha256Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Sha256PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Sha256Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Sha256PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Sha256Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
