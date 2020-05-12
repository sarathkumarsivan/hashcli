[![Build Status](https://travis-ci.org/sarathkumarsivan/hashutils.svg?branch=master)](https://travis-ci.org/sarathkumarsivan/hashutils)
[![codecov](https://codecov.io/gh/sarathkumarsivan/hashutils/branch/master/graph/badge.svg)](https://codecov.io/gh/sarathkumarsivan/hashutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/sarathkumarsivan/hashutils)](https://goreportcard.com/report/github.com/sarathkumarsivan/hashutils)
[![GoDoc](https://godoc.org/github.com/sarathkumarsivan/hashutils?status.svg)](https://godoc.org/github.com/sarathkumarsivan/hashutils)
[![GitHub issues](https://img.shields.io/github/issues/sarathkumarsivan/hashutils)](https://github.com/sarathkumarsivan/hashutils/issues)
[![GitHub](https://img.shields.io/github/license/sarathkumarsivan/hashutils)](https://github.com/sarathkumarsivan/hashutils/blob/master/LICENSE)
# hashutils

This is a collection of hashing utility functions writen in Go.

## Installation
```bash
go get -u github.com/sarathkumarsivan/hashutils
```

## Usage
```go
// MD5 Hash with hexadecimal encoding
hash, _ := MD5Hex("foo")
```