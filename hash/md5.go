package hash

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func Md5(text string) ([]byte, error) {
	hash := md5.New()
	return hashText(hash, text)
}

func MD5Hex(text string) (string, error) {
	hash, err := Md5(text)
	return hex.EncodeToString(hash), err
}

func MD5Base64StdEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.StdEncoding.EncodeToString(hash), err
}

func MD5Base64URLEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.URLEncoding.EncodeToString(hash), err
}

func MD5Base64RawURLEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func MD5Base64RawStdEnc(text string) (string, error) {
	hash, err := Md5(text)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func MD5File(path string) ([]byte, error) {
	hash := md5.New()
	return hashFile(hash, path)
}

func MD5FileHex(path string) (string, error) {
	hash, err := MD5File(path)
	return hex.EncodeToString(hash), err
}

func MD5FileBase64StdEnc(path string) (string, error) {
	hash, err := MD5File(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func MD5FileBase64URLEnc(path string) (string, error) {
	hash, err := MD5File(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func MD5FileBase64RawURLEnc(path string) (string, error) {
	hash, err := MD5File(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func MD5FileBase64RawStdEnc(path string) (string, error) {
	hash, err := MD5File(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func MD5Dir(path string) ([]byte, error) {
	hash := md5.New()
	return hashDir(hash, path)
}

func MD5DirHex(path string) (string, error) {
	hash, err := MD5Dir(path)
	return hex.EncodeToString(hash), err
}

func MD5DirBase64StdEnc(path string) (string, error) {
	hash, err := MD5Dir(path)
	return base64.StdEncoding.EncodeToString(hash), err
}

func MD5DirBase64URLEnc(path string) (string, error) {
	hash, err := MD5Dir(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func MD5DirBase64RawURLEnc(path string) (string, error) {
	hash, err := MD5Dir(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func MD5DirBase64RawStdEnc(path string) (string, error) {
	hash, err := MD5Dir(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}

func MD5Path(path string) ([]byte, error) {
	hash := md5.New()
	return hashPath(hash, path)
}

func MD5PathHex(path string) (string, error) {
	hash, err := MD5Path(path)
	return hex.EncodeToString(hash), err
}

func MD5PathBase64StdEnc(path string) (string, error) {
	hash, err := MD5Path(path)
	return hex.EncodeToString(hash), err
}

func MD5PathBase64URLEnc(path string) (string, error) {
	hash, err := MD5Path(path)
	return base64.URLEncoding.EncodeToString(hash), err
}

func MD5PathBase64RawURLEnc(path string) (string, error) {
	hash, err := MD5Path(path)
	return base64.RawURLEncoding.EncodeToString(hash), err
}

func MD5PathBase64RawStdEnc(path string) (string, error) {
	hash, err := MD5Path(path)
	return base64.RawStdEncoding.EncodeToString(hash), err
}
