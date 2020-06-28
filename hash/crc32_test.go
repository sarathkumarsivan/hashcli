package hash

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCrc32cHashHex(t *testing.T) {
	crc32, err := Crc32HashHex("foo")
	require.NoError(t, err, "Error hashing text to using %s", Crc32Hash)
	assert.Equal(t, "cfc4ae1d", crc32)
}
