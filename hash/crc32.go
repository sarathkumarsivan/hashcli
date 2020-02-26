package hash

import (
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"
)

// Crc32 returns CRC32 checksum of a text as bytes.
func Crc32(text string) []byte {
	table := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(table)
	hash.Write([]byte(text))
	return hash.Sum(nil)[:]
}

// Crc32Hex returns the CRC32 checksum of a text in
// hexadecimal encoding format.
func Crc32Hex(text string) string {
	return hex.EncodeToString(Crc32(text))
}

// Crc32Base64StdEnc returns the CRC32 checksum of a text in
// standard base64 encoding, as defined in RFC 4648.
func Crc32Base64StdEnc(text string) string {
	return base64.StdEncoding.EncodeToString(Crc32(text))
}

// Crc32Base64URLEnc returns the CRC32 checksum of a text in
// an alternate base64 encoding defined in RFC 4648.
func Crc32Base64URLEnc(text string) string {
	return base64.URLEncoding.EncodeToString(Crc32(text))
}

// Crc32Base64RawURLEnc returns the CRC32 checksum of a text in
// a padded alternate base64 encoding defined in RFC 4648.
func Crc32Base64RawURLEnc(text string) string {
	return base64.RawURLEncoding.EncodeToString(Crc32(text))
}

// Crc32Base64RawStdEnc returns the CRC32 checksum of a text in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Crc32Base64RawStdEnc(text string) string {
	return base64.RawStdEncoding.EncodeToString(Crc32(text))
}

// Crc32File returns CRC32 checksum of a file as bytes.
func Crc32File(path string) ([]byte, error) {
	table := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(table)
	return hashFile(hash, path)
}

// Crc32FileHex returns the CRC32 checksum of a file in
// hexadecimal encoding format.
func Crc32FileHex(path string) (string, error) {
	hash, err := Crc32File(path)
	return hex.EncodeToString(hash), err
}

// Crc32FileBase64StdEnc returns the CRC32 checksum of a file in
// standard base64 encoding, as defined in RFC 4648.
func Crc32FileBase64StdEnc(path string) (string, error) {
	hash, err := Crc32File(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}

// Crc32FileBase64URLEnc returns the CRC32 checksum of a file in
// an alternate base64 encoding defined in RFC 4648.
func Crc32FileBase64URLEnc(path string) (string, error) {
	hash, err := Crc32File(path)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(hash), err
}

// Crc32FileBase64RawURLEnc returns the CRC32 checksum of a file in
// a padded alternate base64 encoding defined in RFC 4648.
func Crc32FileBase64RawURLEnc(path string) (string, error) {
	hash, err := Crc32File(path)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Crc32FileBase64RawStdEnc returns the CRC32 checksum of a file in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Crc32FileBase64RawStdEnc(path string) (string, error) {
	hash, err := Crc32File(path)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Crc32Dir returns CRC32 checksum of a directory as bytes.
func Crc32Dir(path string) ([]byte, error) {
	table := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(table)
	return hashDir(hash, path)
}

// Crc32DirHex returns the CRC32 checksum of a directory in
// hexadecimal encoding format.
func Crc32DirHex(path string) (string, error) {
	hash, err := Crc32Dir(path)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), err
}

// Crc32DirBase64StdEnc returns the CRC32 checksum of a directory in
// standard base64 encoding, as defined in RFC 4648.
func Crc32DirBase64StdEnc(path string) (string, error) {
	hash, err := Crc32Dir(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), err
}

// Crc32DirBase64URLEnc returns the CRC32 checksum of a directory in
// an alternate base64 encoding defined in RFC 4648.
func Crc32DirBase64URLEnc(path string) (string, error) {
	hash, err := Crc32Dir(path)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(hash), err
}

// Crc32DirBase64RawURLEnc returns the CRC32 checksum of a directory in
// a padded alternate base64 encoding defined in RFC 4648.
func Crc32DirBase64RawURLEnc(path string) (string, error) {
	hash, err := Crc32Dir(path)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(hash), err
}

// Crc32DirBase64RawStdEnc returns the CRC32 checksum of a directory in
// a standard raw, un-padded base64 encoding, as defined in RFC 4648.
func Crc32DirBase64RawStdEnc(path string) (string, error) {
	hash, err := Crc32Dir(path)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(hash), err
}

// Crc32Path returns CRC32 checksum of a path as bytes.
func Crc32Path(path string) ([]byte, error) {
	table := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(table)
	return hashPath(hash, path)
}

// Crc32PathHex returns the CRC32 checksum of a path in
// hexadecimal encoding format.
func Crc32PathHex(path string) (string, error) {
	hash, err := Crc32Path(path)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), err
}
