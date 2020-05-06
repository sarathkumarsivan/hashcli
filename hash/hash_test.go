package hash

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashText(t *testing.T) {
	maker := New().Algorithm(MD5Hash).Encoding(Hex).Build()
	hash, err := maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash)

	maker = New().Algorithm(SHA1Hash).Encoding(Hex).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", hash)

	maker = New().Algorithm(SHA256Hash).Encoding(Hex).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", hash)

	maker = New().Algorithm(SHA512Hash).Encoding(Hex).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7", hash)

	maker = New().Algorithm(MD5Hash).Encoding(Base64).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", Base64)
	assert.Equal(t, "rL0Y20zC+Fzt72VPzMSk2A==", hash)

	maker = New().Algorithm(SHA1Hash).Encoding(Base64).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", Base64)
	assert.Equal(t, "C+7Hteo/D9vJXQ3UfzxbwnXaijM=", hash)

	maker = New().Algorithm(SHA256Hash).Encoding(Base64).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", Base64)
	assert.Equal(t, "LCa0a2j/xo/5m0U8HTBBNBNCLXBkg7+g+YpeiGJm564=", hash)

	maker = New().Algorithm(SHA512Hash).Encoding(Base64).Build()
	hash, err = maker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", Base64)
	assert.Equal(t, "9/u6bgY2+JDlb7vzKD5STG+jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w==", hash)
}

func TestHashFile(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	defer os.Remove(foo.Name())

	maker := New().Algorithm(MD5Hash).Encoding(Hex).Build()
	hash, err := maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", hash)

	maker = New().Algorithm(SHA1Hash).Encoding(Hex).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", hash)

	maker = New().Algorithm(SHA256Hash).Encoding(Hex).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", hash)

	maker = New().Algorithm(SHA512Hash).Encoding(Hex).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", hash)

	maker = New().Algorithm(MD5Hash).Encoding(Base64).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "1B2M2Y8AsgTpgAmY7PhCfg==", hash)

	maker = New().Algorithm(SHA1Hash).Encoding(Base64).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "2jmj7l5rSw0yVb/vlWAYkK/YBwk=", hash)

	maker = New().Algorithm(SHA256Hash).Encoding(Base64).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=", hash)

	maker = New().Algorithm(SHA512Hash).Encoding(Base64).Build()
	hash, err = maker.HashFile(foo.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==", hash)
}

func TestHashFiles(t *testing.T) {
	foo, err := ioutil.TempFile("", "foo.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = foo.WriteString("foo")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(foo.Name())

	bar, err := ioutil.TempFile("", "bar.*")
	require.NoError(t, err, "Error creating temporary file")
	_, err = bar.WriteString("bar")
	require.NoError(t, err, "Error writing to temporary file")
	defer os.Remove(bar.Name())

	maker := New().Algorithm(MD5Hash).Encoding(Hex).Build()
	hash, err := maker.HashFiles(foo.Name(), bar.Name())
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash[foo.Name()])
	assert.Equal(t, "37b51d194a7513e45b56f6524f2d51f2", hash[bar.Name()])

	maker = New().Algorithm(SHA1Hash).Encoding(Hex).Build()
	hash, err = maker.HashFiles(foo.Name(), bar.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", hash[foo.Name()])
	assert.Equal(t, "62cdb7020ff920e5aa642c3d4066950dd1f01f4d", hash[bar.Name()])

	maker = New().Algorithm(SHA256Hash).Encoding(Hex).Build()
	hash, err = maker.HashFiles(foo.Name(), bar.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", hash[foo.Name()])
	assert.Equal(t, "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9", hash[bar.Name()])

	maker = New().Algorithm(SHA512Hash).Encoding(Hex).Build()
	hash, err = maker.HashFiles(foo.Name(), bar.Name())
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7", hash[foo.Name()])
	assert.Equal(t, "d82c4eb5261cb9c8aa9855edd67d1bd10482f41529858d925094d173fa662aa91ff39bc5b188615273484021dfb16fd8284cf684ccf0fc795be3aa2fc1e6c181", hash[bar.Name()])

}
