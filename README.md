[![Build Status](https://travis-ci.org/sarathkumarsivan/hashutils.svg?branch=master)](https://travis-ci.org/sarathkumarsivan/hashutils)
[![codecov](https://codecov.io/gh/sarathkumarsivan/hashutils/branch/master/graph/badge.svg)](https://codecov.io/gh/sarathkumarsivan/hashutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/sarathkumarsivan/hashutils)](https://goreportcard.com/report/github.com/sarathkumarsivan/hashutils)
[![GoDoc](https://godoc.org/github.com/sarathkumarsivan/hashutils?status.svg)](https://godoc.org/github.com/sarathkumarsivan/hashutils)
[![GitHub issues](https://img.shields.io/github/issues/sarathkumarsivan/hashutils)](https://github.com/sarathkumarsivan/hashutils/issues)
[![GitHub](https://img.shields.io/github/license/sarathkumarsivan/hashutils)](https://github.com/sarathkumarsivan/hashutils/blob/master/LICENSE)
# hashutils

This is a collection of hashing utility functions written in Go.

## Installation
```bash
go get -u github.com/sarathkumarsivan/hashutils
```

## Usage

### MD5 Hash
```scala
// MD5 Hash with hexadecimal encoding.
hash, _ := MD5Hex("foo")

// MD5 Hash with standard base64 encoding.
hash, _ := MD5Base64StdEnc("foo")

// MD5 Hash with alternate base64 encoding.
hash, _ := MD5Base64URLEnc("foo")

// MD5 Hash with standard raw, unpadded base64 encoding.
hash, _ := MD5Base64RawURLEnc("foo")

// MD5 Hash with unpadded alternate base64 encoding.
hash, _ := MD5Base64RawStdEnc("foo")

// MD5 Hash of file with hexadecimal encoding.
hash, _ := MD5FileHex("foo.txt")

// MD5 Hash of file with standard base64 encoding.
hash, _ := MD5FileBase64StdEnc("foo.txt")

// MD5 Hash of file with alternate base64 encoding.
hash, _ := MD5FileBase64URLEnc("foo.txt")

// MD5 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := MD5FileBase64RawURLEnc("foo.txt")

// MD5 Hash of file with unpadded alternate base64 encoding.
hash, _ := MD5FileBase64RawStdEnc("foo.txt")
```

### SHA1 Hash
```scala
// SHA1 Hash with hexadecimal encoding.
hash, _ := SHA1Hex("foo")

// SHA1 Hash with standard base64 encoding.
hash, _ := SHA1Base64StdEnc("foo")

// SHA1 Hash with alternate base64 encoding.
hash, _ := SHA1Base64URLEnc("foo")

// SHA1 Hash with standard raw, unpadded base64 encoding.
hash, _ := SHA1Base64RawURLEnc("foo")

// SHA1 Hash with unpadded alternate base64 encoding.
hash, _ := SHA1Base64RawStdEnc("foo")
```

### SHA224 Hash
```scala
// SHA224 Hash with hexadecimal encoding.
hash, _ := SHA224Hex("foo")

// SHA224 Hash with standard base64 encoding.
hash, _ := SHA224Base64StdEnc("foo")

// SHA224 Hash with alternate base64 encoding.
hash, _ := SHA224Base64URLEnc("foo")

// SHA224 Hash with standard raw, unpadded base64 encoding.
hash, _ := SHA224Base64RawURLEnc("foo")

// SHA224 Hash with unpadded alternate base64 encoding.
hash, _ := SHA224Base64RawStdEnc("foo")
```

### SHA256 Hash
```scala
// SHA256 Hash with hexadecimal encoding.
hash, _ := SHA256Hex("foo")

// SHA256 Hash with standard base64 encoding.
hash, _ := SHA256Base64StdEnc("foo")

// SHA256 Hash with alternate base64 encoding.
hash, _ := SHA256Base64URLEnc("foo")

// SHA256 Hash with standard raw, unpadded base64 encoding.
hash, _ := SHA256Base64RawURLEnc("foo")

// SHA256 Hash with unpadded alternate base64 encoding.
hash, _ := SHA256Base64RawStdEnc("foo")
```

### SHA384 Hash
```scala
// SHA384 Hash with hexadecimal encoding.
hash, _ := SHA384Hex("foo")

// SHA384 Hash with standard base64 encoding.
hash, _ := SHA384Base64StdEnc("foo")

// SHA384 Hash with alternate base64 encoding.
hash, _ := SHA384Base64URLEnc("foo")

// SHA384 Hash with standard raw, unpadded base64 encoding.
hash, _ := SHA384Base64RawURLEnc("foo")

// SHA384 Hash with unpadded alternate base64 encoding.
hash, _ := SHA384Base64RawStdEnc("foo")
```

### SHA512 Hash
```scala
// SHA512 Hash with hexadecimal encoding.
hash, _ := SHA512Hex("foo")

// SHA512 Hash with standard base64 encoding.
hash, _ := SHA512Base64StdEnc("foo")

// SHA512 Hash with alternate base64 encoding.
hash, _ := SHA512Base64URLEnc("foo")

// SHA512 Hash with standard raw, unpadded base64 encoding.
hash, _ := SHA512Base64RawURLEnc("foo")

// SHA512 Hash with unpadded alternate base64 encoding.
hash, _ := SHA512Base64RawStdEnc("foo")
```
