package hashcli

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSHA1Hash(t *testing.T) {
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", GetSHA1Hash("foo"))
	assert.Equal(t, "62cdb7020ff920e5aa642c3d4066950dd1f01f4d", GetSHA1Hash("bar"))
	log.Printf("%v", GetSHA1Hash("bar"))
}
