package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA512Hash(t *testing.T) {
	hash, err := SHA512Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7", hash)

	hash, err = SHA512Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9/u6bgY2+JDlb7vzKD5STG+jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w==", hash)

	hash, err = SHA512Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9/u6bgY2+JDlb7vzKD5STG+jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w", hash)

	hash, err = SHA512Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9_u6bgY2-JDlb7vzKD5STG-jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w", hash)

	hash, err = SHA512Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9_u6bgY2-JDlb7vzKD5STG-jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w==", hash)
}

func TestSHA512HashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer func() { _ = os.Remove(foo.Name()) }()

	hash, err := SHA512FileHex(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", hash)

	hash, err = SHA512FileBase64StdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==", hash)

	hash, err = SHA512FileBase64URLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg_SpIdNs6c5H0NE8XYXysP-DGNKHfuwvY7kxvUdBeoGlODJ6-SfaPg==", hash)

	hash, err = SHA512FileBase64RawURLEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg_SpIdNs6c5H0NE8XYXysP-DGNKHfuwvY7kxvUdBeoGlODJ6-SfaPg", hash)

	hash, err = SHA512FileBase64RawStdEnc(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg", hash)
}

func TestSHA512HashDir(t *testing.T) {
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

	hash, err := SHA512DirHex(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA512Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA512DirBase64StdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA512Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA512DirBase64URLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA512Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA512DirBase64RawURLEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA512Hash)
	assert.NotEmpty(t, hash)

	hash, err = SHA512DirBase64RawStdEnc(dir)
	require.NoError(t, err, "Error hashing dir to using %s", SHA512Hash)
	assert.NotEmpty(t, hash)
}

func TestSHA512HashPath(t *testing.T) {
	dir, err := ioutil.TempDir("", "qux")
	require.NoError(t, err, "Error creating temporary directory")
	defer os.Remove(dir)

	foo, err := ioutil.TempFile(dir, "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())
}
