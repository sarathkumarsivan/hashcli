package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCrc32cHashHex(t *testing.T) {
	assert.Equal(t, "cfc4ae1d", Crc32Hex("foo"))
	assert.Equal(t, "z8SuHQ==", Crc32Base64StdEnc("foo"))
	assert.Equal(t, "z8SuHQ==", Crc32Base64URLEnc("foo"))
	assert.Equal(t, "z8SuHQ", Crc32Base64RawURLEnc("foo"))
	assert.Equal(t, "z8SuHQ", Crc32Base64RawStdEnc("foo"))
}

func TestCrc32HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := Crc32FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Crc32Hash)
	assert.Equal(t, "00000000", hash)

	hash, err = Crc32FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Crc32Hash)
	assert.Equal(t, "AAAAAA==", hash)

	hash, err = Crc32FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Crc32Hash)
	assert.Equal(t, "AAAAAA==", hash)

	hash, err = Crc32FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", Crc32Hash)
	assert.Equal(t, "AAAAAA", hash)
}
