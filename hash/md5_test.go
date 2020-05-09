package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMd5Hash(t *testing.T) {
	hash, err := md5Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash)

	hash, err = md5Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A==", hash)

	hash, err = md5Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A", hash)

	hash, err = md5Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC-Fzt72VPzMSk2A", hash)

	hash, err = md5Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC-Fzt72VPzMSk2A==", hash)
}
