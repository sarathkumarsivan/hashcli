package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA384Hash(t *testing.T) {
	hash, err := SHA384Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA384Hash)
	assert.Equal(t, "98c11ffdfdd540676b1a137cb1a22b2a70350c9a44171d6b1180c6be5cbb2ee3f79d532c8a1dd9ef2e8e08e752a3babb", hash)

	hash, err = SHA384Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "mMEf/f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)

	hash, err = SHA384Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "mMEf/f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)

	hash, err = SHA384Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "mMEf_f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)

	hash, err = SHA384Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "mMEf_f3VQGdrGhN8saIrKnA1DJpEFx1rEYDGvly7LuP3nVMsih3Z7y6OCOdSo7q7", hash)
}
