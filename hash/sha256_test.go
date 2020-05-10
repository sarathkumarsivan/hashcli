package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA256Hash(t *testing.T) {
	hash, err := SHA256Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", hash)

	hash, err = SHA256Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "LCa0a2j/xo/5m0U8HTBBNBNCLXBkg7+g+YpeiGJm564=", hash)

	hash, err = SHA256Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "LCa0a2j/xo/5m0U8HTBBNBNCLXBkg7+g+YpeiGJm564", hash)

	hash, err = SHA256Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564", hash)

	hash, err = SHA256Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564=", hash)
}
