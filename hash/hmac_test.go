package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHMAC(t *testing.T) {
	assert.Equal(t, "147933218aaabc0b8b10a2b3a5c34684c8d94341bcf10a4736dc7270f7741851", HMAC256Hex("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE=", HMAC256Base64StdEnc("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE", HMAC256Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE=", HMAC256Base64URLEnc("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE", HMAC256Base64RawURLEnc("foo", "bar"))

	assert.Equal(t, "d7f508375f4b5b1c236d2df1b850de2474a913644876705e62bd78cc", HMAC224Hex("foo", "bar"))
	assert.Equal(t, "1/UIN19LWxwjbS3xuFDeJHSpE2RIdnBeYr14zA==", HMAC224Base64StdEnc("foo", "bar"))
	assert.Equal(t, "1/UIN19LWxwjbS3xuFDeJHSpE2RIdnBeYr14zA", HMAC224Base64RawStdEnc("foo", "bar"))
}
