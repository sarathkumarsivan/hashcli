package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA1Hash(t *testing.T) {
	hash, err := Sha1Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", hash)

	hash, err = Sha1Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C+7Hteo/D9vJXQ3UfzxbwnXaijM=", hash)

	hash, err = Sha1Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C+7Hteo/D9vJXQ3UfzxbwnXaijM", hash)

	hash, err = Sha1Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C-7Hteo_D9vJXQ3UfzxbwnXaijM", hash)

	hash, err = Sha1Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "C-7Hteo_D9vJXQ3UfzxbwnXaijM=", hash)
}

func TestSHA1HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := Sha1FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", hash)

	hash, err = Sha1FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "2jmj7l5rSw0yVb/vlWAYkK/YBwk=", hash)

	hash, err = Sha1FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "2jmj7l5rSw0yVb_vlWAYkK_YBwk=", hash)

	hash, err = Sha1FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "2jmj7l5rSw0yVb_vlWAYkK_YBwk", hash)

	hash, err = Sha1FileBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "2jmj7l5rSw0yVb/vlWAYkK/YBwk", hash)
}

func TestSHA1HashDir(t *testing.T) {
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

	hash, err := Sha1DirHex(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1DirBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1DirBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1DirBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1DirBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)
}

func TestSHA1HashPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	hash, err := Sha1PathHex(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1PathHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1PathBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1PathBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1PathBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = Sha1PathBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA1PathBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA1PathBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA1PathBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA1PathBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.NotEmpty(t, hash)
}
