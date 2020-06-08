package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func SHA384(text string) ([]byte, error) {
	hash := sha512.New384()
	return hashText(hash, text)
}

func SHA384Hex(text string) (string, error) {
	hash, err := SHA384(text)
	return hex.EncodeToString(hash), err
}

func SHA384Base64StdEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA384Base64URLEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA384Base64RawURLEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA384Base64RawStdEnc(text string) (string, error) {
	hash, err := SHA384(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA384File(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashFile(hash, path)
}

func SHA384FileHex(path string) (string, error) {
	hash, err := SHA384File(path)
	return hex.EncodeToString(hash), err
}

func SHA384FileBase64StdEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA384FileBase64URLEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA384FileBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA384FileBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA384File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA384Dir(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashDir(hash, path)
}

func SHA384DirHex(path string) (string, error) {
	hash, err := SHA384Dir(path)
	return hex.EncodeToString(hash), err
}

func SHA384DirBase64StdEnc(path string) (string, error) {
	hash, err := SHA384Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func SHA384DirBase64URLEnc(path string) (string, error) {
	hash, err := SHA384Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func SHA384DirBase64RawURLEnc(path string) (string, error) {
	hash, err := SHA384Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func SHA384DirBase64RawStdEnc(path string) (string, error) {
	hash, err := SHA384Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func SHA384Path(path string) ([]byte, error) {
	hash := sha512.New384()
	return hashPath(hash, path)
}

func SHA384PathHex(path string) (string, error) {
	hash, err := SHA384Path(path)
	return hex.EncodeToString(hash), err
}

func SHA384PathBase64StdEnc(path string) (string, error) {
	hash, err := SHA384Path(path)
	return hex.EncodeToString(hash), err
}
