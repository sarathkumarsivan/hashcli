package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFnvHash(t *testing.T) {
	fnv32, err := fnv32("foo")
	require.NoError(t, err, "Error hashing text to using %s", FNV32)
	assert.Equal(t, uint32(0x408f5e13), fnv32)

	fnv32, err = fnv32a("foo")
	require.NoError(t, err, "Error hashing text to using %s", FNV32a)
	assert.Equal(t, uint32(0xa9f37ed7), fnv32)

	fnv64, err := fnv64("foo")
	require.NoError(t, err, "Error hashing text to using %s", FNV64)
	assert.Equal(t, uint64(0xd8cbc7186ba13533), fnv64)

	fnv64, err = fnv64a("foo")
	require.NoError(t, err, "Error hashing text to using %s", FNV64a)
	assert.Equal(t, uint64(0xdcb27518fed9d577), fnv64)
}
