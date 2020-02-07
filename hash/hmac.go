package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// HmacMd5 computes a Hash-based Message Authentication Code (HMAC) by using
// the MD5 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacMd5(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(md5.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

// HmacMd5Hex mixes a secret key with the message data, hashes the result with the
// MD5 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using hexadecimal encoding.
func HmacMd5Hex(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return hex.EncodeToString(bytes)
}

// HmacMd5Base64StdEnc mixes a secret key with the message data, hashes the result with
// the MD5 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 standard encoding.
func HmacMd5Base64StdEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

// HmacMd5Base64RawStdEnc mixes a secret key with the message data, hashes the result with
// the MD5 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using raw base64 standard encoding.
func HmacMd5Base64RawStdEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

// HmacMd5Base64URLEnc mixes a secret key with the message data, hashes the result with
// the MD5 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 URL encoding.
func HmacMd5Base64URLEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

// HmacMd5Base64RawURLEnc mixes a secret key with the message data, hashes the result with
// the MD5 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 raw URL encoding.
func HmacMd5Base64RawURLEnc(message string, secret string) string {
	bytes := HmacMd5(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// HmacSha1 computes a Hash-based Message Authentication Code (HMAC) by using
// the SHA-1 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacSha1(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

// HmacSha1Hex mixes a secret key with the message data, hashes the result with the
// SHA-1 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using hexadecimal encoding.
func HmacSha1Hex(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return hex.EncodeToString(bytes)
}

// HmacSha1Base64StdEnc mixes a secret key with the message data, hashes the result with
// the SHA-1 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 standard encoding.
func HmacSha1Base64StdEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

// HmacSha1Base64RawStdEnc mixes a secret key with the message data, hashes the result with
// the SHA-1 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using raw base64 standard encoding.
func HmacSha1Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

// HmacSha1Base64URLEnc mixes a secret key with the message data, hashes the result with
// the SHA-1 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 URL encoding.
func HmacSha1Base64URLEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

// HmacSha1Base64RawURLEnc mixes a secret key with the message data, hashes the result with
// the SHA-1 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 raw URL encoding.
func HmacSha1Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha1(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// HmacSha256 computes a Hash-based Message Authentication Code (HMAC) by using
// the SHA-256 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacSha256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

// HmacSha256Hex mixes a secret key with the message data, hashes the result with the
// SHA-256 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using hexadecimal encoding.
func HmacSha256Hex(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return hex.EncodeToString(bytes)
}

// HmacSha256Base64StdEnc mixes a secret key with the message data, hashes the result with
// the SHA-256 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 standard encoding.
func HmacSha256Base64StdEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

// HmacSha256Base64RawStdEnc mixes a secret key with the message data, hashes the result with
// the SHA-256 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using raw base64 standard encoding.
func HmacSha256Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

// HmacSha256Base64URLEnc mixes a secret key with the message data, hashes the result with
// the SHA-256 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using raw base64 standard encoding.
func HmacSha256Base64URLEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

// HmacSha256Base64RawURLEnc mixes a secret key with the message data, hashes the result with
// the SHA-256 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 raw URL encoding.
func HmacSha256Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha256(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// HmacSha224 computes a Hash-based Message Authentication Code (HMAC) by using
// the SHA-224 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacSha224(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New224, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

// HmacSha224Hex mixes a secret key with the message data, hashes the result with the
// SHA-224 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using hexadecimal encoding.
func HmacSha224Hex(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return hex.EncodeToString(bytes)
}

// HmacSha224Base64StdEnc mixes a secret key with the message data, hashes the result with
// the SHA-224 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 standard encoding.
func HmacSha224Base64StdEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

// HmacSha224Base64RawStdEnc mixes a secret key with the message data, hashes the result with
// the SHA-224 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using raw base64 standard encoding.
func HmacSha224Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

// HmacSha224Base64URLEnc mixes a secret key with the message data, hashes the result with
// the SHA-224 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using raw base64 standard encoding.
func HmacSha224Base64URLEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

// HmacSha224Base64RawURLEnc mixes a secret key with the message data, hashes the result with
// the SHA-224 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using base64 raw URL encoding.
func HmacSha224Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha224(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// HmacSha512 computes a Hash-based Message Authentication Code (HMAC) by using
// the SHA-512 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacSha512(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

// HmacSha512Hex mixes a secret key with the message data, hashes the result with the
// SHA-512 hash function, mixes that hash value with the secret key again, then applies
// the hash function a second time, and encodes the result using hexadecimal encoding.
func HmacSha512Hex(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha512Base64StdEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha512Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha512Base64URLEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha512Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha512(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// HmacSha384 computes a Hash-based Message Authentication Code (HMAC) by using
// the SHA-384 hash function. The HMAC process mixes a secret key with the message
// data, hashes the result with the hash function, mixes that hash value with
// the secret key again, and then applies the hash function a second time.
func HmacSha384(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha512.New384, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func HmacSha384Hex(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return hex.EncodeToString(bytes)
}

func HmacSha384Base64StdEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HmacSha384Base64RawStdEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func HmacSha384Base64URLEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.URLEncoding.EncodeToString(bytes)
}

func HmacSha384Base64RawURLEnc(message string, secret string) string {
	bytes := HmacSha384(message, secret)
	return base64.RawURLEncoding.EncodeToString(bytes)
}
