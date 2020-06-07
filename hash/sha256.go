package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SHA256(text string) ([]byte, error) {
	hash := sha256.New()
	return hashText(hash, text)
}

func SHA256Hex(text string) (string, error) {
	hash, err := SHA256(text)
	return hex.EncodeToString(hash), err
}

func SHA256Base64StdEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA256Base64URLEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA256(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA256File(path string) ([]byte, error) {
	hash := sha256.New()
	return hashFile(hash, path)
}

func SHA256FileHex(path string) (string, error) {
	hash, err := SHA256File(path)
	return hex.EncodeToString(hash), err
}

func SHA256FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA256FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA256File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA256Dir(path string) ([]byte, error) {
	hash := sha256.New()
	return hashDir(hash, path)
}

func SHA256DirHex(path string) (string, error) {
	hash, err := SHA256Dir(path)
	return hex.EncodeToString(hash), err
}

func SHA256DirBase64StdEnc(path string) (string, error) {
	hash, err := SHA256Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA256DirBase64URLEnc(path string) (string, error) {
	hash, err := SHA256Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA256DirBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA256Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA256DirBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA256Dir(path)
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
