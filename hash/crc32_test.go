package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrc32cHashHex(t *testing.T) {
	assert.Equal(t, "cfc4ae1d", Crc32Hex("foo"))
}
