package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMD5Hash(t *testing.T) {
	hash, err := MD5Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash)

	hash, err = MD5Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A==", hash)

	hash, err = MD5Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A", hash)

	hash, err = MD5Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC-Fzt72VPzMSk2A", hash)

	hash, err = MD5Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "rL0Y20zC-Fzt72VPzMSk2A==", hash)
}

func TestMD5HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := MD5FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", hash)

	hash, err = MD5FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg==", hash)

	hash, err = MD5FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg==", hash)

	hash, err = MD5FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg==", hash)
}
