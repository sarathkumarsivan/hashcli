package hash

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// Md5 returns MD5 checksum of a text as bytes.
func Md5(text string) ([]byte, error) {
	hash := md5.New()
	return hashText(hash, text)
}

// Md5Hex returns the MD5 checksum of a text in
// hexadecimal encoding format.
func Md5Hex(text string) (string, error) {
	hash, err := Md5(text)
	return hex.EncodeToString(hash), err
}

// Md5Base64StdEnc returns the MD5 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Md5Base64StdEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Md5Base64URLEnc returns the MD5 checksum of a text in
// an alternate base64 encoding defined in RFC 4648.
func Md5Base64URLEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Md5Base64RawURLEnc returns the MD5 checksum of a text in
// a padded alternate base64 encoding defined in RFC 4648.
func Md5Base64RawURLEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Md5Base64RawStdEnc returns the MD5 checksum of a text in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Md5Base64RawStdEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Md5File returns MD5 checksum of a file as bytes.
func Md5File(path string) ([]byte, error) {
	hash := md5.New()
	return hashFile(hash, path)
}

// Md5FileHex returns the MD5 checksum of a file in
// hexadecimal encoding format.
func Md5FileHex(path string) (string, error) {
	hash, err := Md5File(path)
	return hex.EncodeToString(hash), err
}

// Md5FileBase64StdEnc returns the MD5 checksum of a file in
// standard base64 encoding, as defined in RFC 4648.
func Md5FileBase64StdEnc(path string) (string, error) {
	hash, err := Md5File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Md5FileBase64URLEnc returns the MD5 checksum of a file in
// an alternate base64 encoding defined in RFC 4648.
func Md5FileBase64URLEnc(path string) (string, error) {
	hash, err := Md5File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Md5FileBase64RawURLEnc returns the MD5 checksum of a file in
// a padded alternate base64 encoding defined in RFC 4648.
func Md5FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Md5File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Md5FileBase64RawStdEnc returns the MD5 checksum of a file in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Md5FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Md5File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Md5Dir returns MD5 checksum of a directory as bytes.
func Md5Dir(path string) ([]byte, error) {
	hash := md5.New()
	return hashDir(hash, path)
}

// Md5DirHex returns the MD5 checksum of a directory in
// hexadecimal encoding format.
func Md5DirHex(path string) (string, error) {
	hash, err := Md5Dir(path)
	return hex.EncodeToString(hash), err
}

// Md5DirBase64StdEnc returns the MD5 checksum of a directory in
// standard base64 encoding, as defined in RFC 4648.
func Md5DirBase64StdEnc(path string) (string, error) {
	hash, err := Md5Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

// Md5DirBase64URLEnc returns the MD5 checksum of a directory in
// an alternate base64 encoding defined in RFC 4648.
func Md5DirBase64URLEnc(path string) (string, error) {
	hash, err := Md5Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

// Md5DirBase64RawURLEnc returns the MD5 checksum of a directory in
// a padded alternate base64 encoding defined in RFC 4648.
func Md5DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Md5Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Md5DirBase64RawStdEnc returns the MD5 checksum of a directory in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Md5DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Md5Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Md5Path returns MD5 checksum of a path as bytes.
func Md5Path(path string) ([]byte, error) {
	hash := md5.New()
	return hashPath(hash, path)
}

// Md5PathHex returns the MD5 checksum of a path in
// hexadecimal encoding format.
func Md5PathHex(path string) (string, error) {
	hash, err := Md5Path(path)
	return hex.EncodeToString(hash), err
}

// Md5PathBase64StdEnc returns the MD5 checksum of a path in
// standard base64 encoding, as defined in RFC 4648.
func Md5PathBase64StdEnc(path string) (string, error) {
	hash, err := Md5Path(path)
	return hex.EncodeToString(hash), err
}

// Md5PathBase64URLEnc returns the MD5 checksum of a path in
// an alternate base64 encoding defined in RFC 4648.
func Md5PathBase64URLEnc(path string) (string, error) {
	hash, err := Md5Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func Md5PathBase64RawURLEnc(path string) (string, error) {
	hash, err := Md5Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func Md5PathBase64RawStdEnc(path string) (string, error) {
	hash, err := Md5Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
