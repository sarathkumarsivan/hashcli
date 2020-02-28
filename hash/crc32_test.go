package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrc32cHashHex(t *testing.T) {
	assert.Equal(t, "cfc4ae1d", Crc32Hex("foo"))
	assert.Equal(t, "z8SuHQ==", Crc32Base64StdEnc("foo"))
	assert.Equal(t, "z8SuHQ==", Crc32Base64URLEnc("foo"))
	assert.Equal(t, "z8SuHQ", Crc32Base64RawURLEnc("foo"))
}
