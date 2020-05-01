package hashcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {
	assert.Equal(t, "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", GetSHA1Hash("foo"))
	assert.Equal(t, "62cdb7020ff920e5aa642c3d4066950dd1f01f4d", GetSHA1Hash("bar"))
	assert.Equal(t, "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", GetSHA256Hash("foo"))
	assert.Equal(t, "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9", GetSHA256Hash("bar"))
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", GetMD5Hash("foo"))
	assert.Equal(t, "acbd18db4cc2f85cedef654fccc4a4d8", GetMD5Hash("foo"))
}
