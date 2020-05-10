package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA384Hash(t *testing.T) {
	hash, err := SHA384Hex("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA384Hash)
	assert.Equal(t, "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7", hash)

	hash, err = SHA384Base64StdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9/u6bgY2+JDlb7vzKD5STG+jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w==", hash)

	hash, err = SHA384Base64RawStdEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9/u6bgY2+JDlb7vzKD5STG+jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w", hash)

	hash, err = SHA384Base64RawURLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9_u6bgY2-JDlb7vzKD5STG-jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w", hash)

	hash, err = SHA384Base64URLEnc("foo")
	require.NoError(t, err, "Error hashing text to using %s", SHA512Hash)
	assert.Equal(t, "9_u6bgY2-JDlb7vzKD5STG-jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w==", hash)
}
