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
hash, _ := Md5Hex("foo")

// MD5 Hash with standard base64 encoding.
hash, _ := Md5Base64StdEnc("foo")

// MD5 Hash with alternate base64 encoding.
hash, _ := Md5Base64URLEnc("foo")

// MD5 Hash with standard raw, unpadded base64 encoding.
hash, _ := Md5Base64RawURLEnc("foo")

// MD5 Hash with unpadded alternate base64 encoding.
hash, _ := Md5Base64RawStdEnc("foo")

// MD5 Hash of file with hexadecimal encoding.
hash, _ := Md5FileHex("foo.txt")

// MD5 Hash of file with standard base64 encoding.
hash, _ := Md5FileBase64StdEnc("foo.txt")

// MD5 Hash of file with alternate base64 encoding.
hash, _ := Md5FileBase64URLEnc("foo.txt")

// MD5 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := Md5FileBase64RawURLEnc("foo.txt")

// MD5 Hash of file with unpadded alternate base64 encoding.
hash, _ := Md5FileBase64RawStdEnc("foo.txt")
```

### SHA1 Hash
```scala
// SHA1 Hash with hexadecimal encoding.
hash, _ := Sha1Hex("foo")

// SHA1 Hash with standard base64 encoding.
hash, _ := Sha1Base64StdEnc("foo")

// SHA1 Hash with alternate base64 encoding.
hash, _ := Sha1Base64URLEnc("foo")

// SHA1 Hash with standard raw, unpadded base64 encoding.
hash, _ := Sha1Base64RawURLEnc("foo")

// SHA1 Hash with unpadded alternate base64 encoding.
hash, _ := Sha1Base64RawStdEnc("foo")

// SHA1 Hash of file with hexadecimal encoding.
hash, _ := Sha1FileHex("foo.txt")

// SHA1 Hash of file with standard base64 encoding.
hash, _ := Sha1FileBase64StdEnc("foo.txt")

// SHA1 Hash of file with alternate base64 encoding.
hash, _ := Sha1FileBase64URLEnc("foo.txt")

// SHA1 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := Sha1FileBase64RawURLEnc("foo.txt")

// SHA1 Hash of file with unpadded alternate base64 encoding.
hash, _ := Sha1FileBase64RawStdEnc("foo.txt")
```

### SHA224 Hash
```scala
// SHA224 Hash with hexadecimal encoding.
hash, _ := Sha224Hex("foo")

// SHA224 Hash with standard base64 encoding.
hash, _ := Sha224Base64StdEnc("foo")

// SHA224 Hash with alternate base64 encoding.
hash, _ := Sha224Base64URLEnc("foo")

// SHA224 Hash with standard raw, unpadded base64 encoding.
hash, _ := Sha224Base64RawURLEnc("foo")

// SHA224 Hash with unpadded alternate base64 encoding.
hash, _ := Sha224Base64RawStdEnc("foo")

// SHA224 Hash of file with hexadecimal encoding.
hash, _ := Sha224FileHex("foo.txt")

// SHA224 Hash of file with standard base64 encoding.
hash, _ := Sha224FileBase64StdEnc("foo.txt")

// SHA224 Hash of file with alternate base64 encoding.
hash, _ := Sha224FileBase64URLEnc("foo.txt")

// SHA224 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := Sha224FileBase64RawURLEnc("foo.txt")

// SHA224 Hash of file with unpadded alternate base64 encoding.
hash, _ := Sha224BFilease64RawStdEnc("foo.txt")
```

### SHA256 Hash
```scala
// SHA256 Hash with hexadecimal encoding.
hash, _ := Sha256Hex("foo")

// SHA256 Hash with standard base64 encoding.
hash, _ := Sha256Base64StdEnc("foo")

// SHA256 Hash with alternate base64 encoding.
hash, _ := Sha256Base64URLEnc("foo")

// SHA256 Hash with standard raw, unpadded base64 encoding.
hash, _ := Sha256Base64RawURLEnc("foo")

// SHA256 Hash with unpadded alternate base64 encoding.
hash, _ := Sha256Base64RawStdEnc("foo")

// SHA256 Hash with hexadecimal encoding.
hash, _ := Sha256FileHex("foo.txt")

// SHA256 Hash of file with standard base64 encoding.
hash, _ := Sha256FileBase64StdEnc("foo.txt")

// SHA256 Hash of file with alternate base64 encoding.
hash, _ := Sha256FileBase64URLEnc("foo.txt")

// SHA256 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := Sha256FileBase64RawURLEnc("foo.txt")

// SHA256 Hash of file with unpadded alternate base64 encoding.
hash, _ := Sha256FileBase64RawStdEnc("foo.txt")
```

### SHA384 Hash
```scala
// SHA384 Hash with hexadecimal encoding.
hash, _ := Sha384Hex("foo")

// SHA384 Hash with standard base64 encoding.
hash, _ := Sha384Base64StdEnc("foo")

// SHA384 Hash with alternate base64 encoding.
hash, _ := Sha384Base64URLEnc("foo")

// SHA384 Hash with standard raw, unpadded base64 encoding.
hash, _ := Sha384Base64RawURLEnc("foo")

// SHA384 Hash with unpadded alternate base64 encoding.
hash, _ := Sha384Base64RawStdEnc("foo")

// SHA384 Hash with hexadecimal encoding.
hash, _ := Sha384FileHex("foo.txt")

// SHA384 Hash of file with standard base64 encoding.
hash, _ := Sha384FileBase64StdEnc("foo.txt")

// SHA384 Hash of file with alternate base64 encoding.
hash, _ := Sha384FileBase64URLEnc("foo.txt")

// SHA384 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := Sha384FileBase64RawURLEnc("foo.txt")

// SHA384 Hash of file with unpadded alternate base64 encoding.
hash, _ := Sha384FileBase64RawStdEnc("foo.txt")
```

### SHA512 Hash
```scala
// SHA512 Hash with hexadecimal encoding.
hash, _ := Sha512Hex("foo")

// SHA512 Hash with standard base64 encoding.
hash, _ := Sha512Base64StdEnc("foo")

// SHA512 Hash with alternate base64 encoding.
hash, _ := Sha512Base64URLEnc("foo")

// SHA512 Hash with standard raw, unpadded base64 encoding.
hash, _ := Sha512Base64RawURLEnc("foo")

// SHA512 Hash with unpadded alternate base64 encoding.
hash, _ := Sha512Base64RawStdEnc("foo")

// SHA512 Hash with hexadecimal encoding.
hash, _ := Sha512FileHex("foo.txt")

// SHA512 Hash of file with standard base64 encoding.
hash, _ := Sha512FileBase64StdEnc("foo.txt")

// SHA512 Hash of file with alternate base64 encoding.
hash, _ := Sha512FileBase64URLEnc("foo.txt")

// SHA512 Hash of file with standard raw, unpadded base64 encoding.
hash, _ := Sha512FileBase64RawURLEnc("foo.txt")

// SHA512 Hash of file with unpadded alternate base64 encoding.
hash, _ := Sha512FileBase64RawStdEnc("foo.txt")

// SHA512 Hash of directory path with hexadecimal encoding
hash, _ := Sha512PathHex("/home/foo")

// SHA512 Hash of file path with hexadecimal encoding
hash, _ := Sha512PathHex("/home/foo.txt")

// SHA512 Hash of directory path with standard base64 encoding.
hash, _ := Sha512PathBase64StdEnc("/home/foo")

// SHA512 Hash of file path with standard base64 encoding.
hash, _ := Sha512PathBase64StdEnc("/home/foo.txt")

// SHA512 Hash of directory path with alternate base64 encoding.
hash, _ := Sha512PathBase64URLEnc("/home/foo")

// SHA512 Hash of file path with alternate base64 encoding.
hash, _ := Sha512PathBase64URLEnc("/home/foo.txt")
```

### Install from source and run through commandline
You can make the hash of a text, file/directory by running the command line tool. 
Install the package from GitHub.
```scala
go install github.com/sarathkumarsivan/hashutils
```
Set the alias in your ~/.profile file and reload the profile as shown below:
```scala
# hashutils Path Settings
alias hash="$GOPATH/bin/hashutils $@"
source ~/.profile
```
Now, you can run hash commands:
```scala
hash -t foo
```
It should display the SHA1 hash of the text "foo" as JSON. SHA1 is used as the default algorithm and results will be
 encoded in base64 format.
```scala
{"text":"foo","algorithm":"sha1","encoding":"hex","hash":"0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"}
```
Try specifying the algorithm and encoding like this:
```scala
hash -t foo -a md5 -e base64
```
You should get the MD5 hash of the text "foo" encoded as "base64"
```scala
{"text":"foo","algorithm":"md5","encoding":"base64","hash":"rL0Y20zC+Fzt72VPzMSk2A=="}
```