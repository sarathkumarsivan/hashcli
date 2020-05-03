package hashcli

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashText(t *testing.T) {
	hashMaker := New().Algorithm(MD5Hash).Encoding(Hex).Build()
	hash, err := hashMaker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash)
	hash, err = hashMaker.HashText("bar")
	require.NoError(t, err, "Error hashing text to using %s", MD5Hash)
	assert.Equal(t, "37b51d194a7513e45b56f6524f2d51f2", hash)

	hashMaker = New().Algorithm(SHA1Hash).Encoding(Hex).Build()
	hash, err = hashMaker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", hash)
	hash, err = hashMaker.HashText("bar")
	require.NoError(t, err, "Error hashing text to using %s", SHA1Hash)
	assert.Equal(t, "62cdb7020ff920e5aa642c3d4066950dd1f01f4d", hash)

	hashMaker = New().Algorithm(SHA256Hash).Encoding(Hex).Build()
	hash, err = hashMaker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", hash)
	hash, err = hashMaker.HashText("bar")
	require.NoError(t, err, "Error hashing text to using %s", SHA256Hash)
	assert.Equal(t, "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9", hash)

	hashMaker = New().Algorithm(SHA512Hash).Encoding(Hex).Build()
	hash, err = hashMaker.HashText("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7", hash)
	hash, err = hashMaker.HashText("bar")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "d82c4eb5261cb9c8aa9855edd67d1bd10482f41529858d925094d173fa662aa91ff39bc5b188615273484021dfb16fd8284cf684ccf0fc795be3aa2fc1e6c181", hash)
}
