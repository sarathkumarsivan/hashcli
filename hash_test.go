package hashcli

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {
	hash, err := GetHash("foo", AlgorithmMD5)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmMD5)
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", hash)
	hash, err = GetHash("bar", AlgorithmMD5)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmMD5)
	assert.Equal(t, "37b51d194a7513e45b56f6524f2d51f2", hash)

	hash, err = GetHash("foo", AlgorithmSHA1)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmSHA1)
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", hash)
	hash, err = GetHash("bar", AlgorithmSHA1)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmSHA1)
	assert.Equal(t, "62cdb7020ff920e5aa642c3d4066950dd1f01f4d", hash)

	hash, err = GetHash("foo", AlgorithmSHA256)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmSHA256)
	assert.Equal(t, "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", hash)
	hash, err = GetHash("bar", AlgorithmSHA256)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmSHA256)
	assert.Equal(t, "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9", hash)

	hash, err = GetHash("foo", AlgorithmSHA512)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmSHA512)
	assert.Equal(t, "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7", hash)
	hash, err = GetHash("bar", AlgorithmSHA512)
	require.NoError(t, err, "Error hashing text to using %s", AlgorithmSHA512)
	assert.Equal(t, "d82c4eb5261cb9c8aa9855edd67d1bd10482f41529858d925094d173fa662aa91ff39bc5b188615273484021dfb16fd8284cf684ccf0fc795be3aa2fc1e6c181", hash)
}
