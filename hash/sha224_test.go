package hash

import (
	"io/ioutil"
	"os"
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

func TestSHA224HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := SHA224FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", hash)

	hash, err = SHA224FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "0UoCjCo6K8lHYQK7KII0xBWisB+CjqYqxbPkLw==", hash)

	hash, err = SHA224FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "0UoCjCo6K8lHYQK7KII0xBWisB-CjqYqxbPkLw==", hash)

	hash, err = SHA224FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "0UoCjCo6K8lHYQK7KII0xBWisB-CjqYqxbPkLw", hash)

	hash, err = SHA224FileBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.Equal(t, "0UoCjCo6K8lHYQK7KII0xBWisB+CjqYqxbPkLw", hash)
}

func TestSHA224HashDir(t *testing.T) {
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

	hash, err := SHA224DirHex(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA224DirBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA224DirBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA224DirBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA224DirBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)
}

func TestSHA22HashPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	hash, err := SHA224PathHex(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA224PathHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA224PathBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing text to using %s", SHA224Hash)
	assert.NotEmpty(t, hash)

}
