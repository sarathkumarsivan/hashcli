package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA1Hash(t *testing.T) {
	hash, err := SHA1Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", hash)

	hash, err = SHA1Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C+7Hteo/D9vJXQ3UfzxbwnXaijM=", hash)

	hash, err = SHA1Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C+7Hteo/D9vJXQ3UfzxbwnXaijM", hash)

	hash, err = SHA1Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C-7Hteo_D9vJXQ3UfzxbwnXaijM", hash)

	hash, err = SHA1Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C-7Hteo_D9vJXQ3UfzxbwnXaijM=", hash)
}
