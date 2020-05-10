package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA224Hash(t *testing.T) {
	hash, err := SHA224Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "0808f64e60d58979fcb676c96ec938270dea42445aeefcd3a4e6f8db", hash)

	hash, err = SHA224Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "CAj2TmDViXn8tnbJbsk4Jw3qQkRa7vzTpOb42w==", hash)

	hash, err = SHA224Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "CAj2TmDViXn8tnbJbsk4Jw3qQkRa7vzTpOb42w", hash)

	hash, err = SHA224Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "CAj2TmDViXn8tnbJbsk4Jw3qQkRa7vzTpOb42w", hash)

	hash, err = SHA224Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "CAj2TmDViXn8tnbJbsk4Jw3qQkRa7vzTpOb42w==", hash)
}
