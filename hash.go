package main

import (
	"crypto/sha1"
	"io"
)

type Hasher interface {
	Hash() string
}

type SHA1Hash struct {
	text string
}

func (s SHA1Hash) Hash() string {
	h := sha1.New()
	io.WriteString(h, s.text)
	return string(h.Sum(nil))
}
