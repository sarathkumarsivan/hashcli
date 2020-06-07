package hash

import (
	"io/ioutil"
	"os"
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

func TestSHA256HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := SHA256FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", hash)

	hash, err = SHA256FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=", hash)

	hash, err = SHA256FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU=", hash)

	hash, err = SHA256FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU", hash)

	hash, err = SHA256FileBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU", hash)
}

func TestSHA256HashDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	bar, err := ioutil.TempFile(dir, "bar.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = bar.WriteString("bar")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(bar.Name())

	hash, err := SHA256DirHex(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA256DirBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA256DirBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA256DirBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA256DirBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)
}

func TestSHA256HashPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	hash, err := SHA256PathHex(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA256PathHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA256PathBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.NotEmpty(t, hash)
}
