package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHMAC(t *testing.T) {
	assert.Equal(t, "31b6db9e5eb4addb42f1a6ca07367adc", HmacMd5Hex("foo", "bar"))
	assert.Equal(t, "Mbbbnl60rdtC8abKBzZ63A==", HmacMd5Base64StdEnc("foo", "bar"))
	assert.Equal(t, "Mbbbnl60rdtC8abKBzZ63A", HmacMd5Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "Mbbbnl60rdtC8abKBzZ63A==", HmacMd5Base64URLEnc("foo", "bar"))
	assert.Equal(t, "Mbbbnl60rdtC8abKBzZ63A", HmacMd5Base64RawURLEnc("foo", "bar"))

	assert.Equal(t, "85d155c55ed286a300bd1cf124de08d87e914f3a", Hmac1Hex("foo", "bar"))
	assert.Equal(t, "hdFVxV7ShqMAvRzxJN4I2H6RTzo=", Hmac1Base64StdEnc("foo", "bar"))
	assert.Equal(t, "hdFVxV7ShqMAvRzxJN4I2H6RTzo", Hmac1Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "hdFVxV7ShqMAvRzxJN4I2H6RTzo=", Hmac1Base64URLEnc("foo", "bar"))
	assert.Equal(t, "hdFVxV7ShqMAvRzxJN4I2H6RTzo", Hmac1Base64RawURLEnc("foo", "bar"))

	assert.Equal(t, "147933218aaabc0b8b10a2b3a5c34684c8d94341bcf10a4736dc7270f7741851", Hmac256Hex("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE=", Hmac256Base64StdEnc("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE", Hmac256Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE=", Hmac256Base64URLEnc("foo", "bar"))
	assert.Equal(t, "FHkzIYqqvAuLEKKzpcNGhMjZQ0G88QpHNtxycPd0GFE", HMAC256Base64RawURLEnc("foo", "bar"))

	assert.Equal(t, "d7f508375f4b5b1c236d2df1b850de2474a913644876705e62bd78cc", HMAC224Hex("foo", "bar"))
	assert.Equal(t, "1/UIN19LWxwjbS3xuFDeJHSpE2RIdnBeYr14zA==", HMAC224Base64StdEnc("foo", "bar"))
	assert.Equal(t, "1/UIN19LWxwjbS3xuFDeJHSpE2RIdnBeYr14zA", HMAC224Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "1_UIN19LWxwjbS3xuFDeJHSpE2RIdnBeYr14zA==", HMAC224Base64URLEnc("foo", "bar"))
	assert.Equal(t, "1_UIN19LWxwjbS3xuFDeJHSpE2RIdnBeYr14zA", HMAC224Base64RawURLEnc("foo", "bar"))

	assert.Equal(t, "24257d7210582a65c731ec55159c8184cc24c02489453e58587f71f44c23a2d61b4b72154a89d17b2d49448a8452ea066f4fc56a2bcead45c088572ffccdb3d8", HMAC512Hex("foo", "bar"))
	assert.Equal(t, "JCV9chBYKmXHMexVFZyBhMwkwCSJRT5YWH9x9EwjotYbS3IVSonRey1JRIqEUuoGb0/FaivOrUXAiFcv/M2z2A==", HMAC512Base64StdEnc("foo", "bar"))
	assert.Equal(t, "JCV9chBYKmXHMexVFZyBhMwkwCSJRT5YWH9x9EwjotYbS3IVSonRey1JRIqEUuoGb0/FaivOrUXAiFcv/M2z2A", HMAC512Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "JCV9chBYKmXHMexVFZyBhMwkwCSJRT5YWH9x9EwjotYbS3IVSonRey1JRIqEUuoGb0_FaivOrUXAiFcv_M2z2A==", HMAC512Base64URLEnc("foo", "bar"))
	assert.Equal(t, "JCV9chBYKmXHMexVFZyBhMwkwCSJRT5YWH9x9EwjotYbS3IVSonRey1JRIqEUuoGb0_FaivOrUXAiFcv_M2z2A", HMAC512Base64RawURLEnc("foo", "bar"))

	assert.Equal(t, "1d9070d07cb7746e0664cccc6cec1fa996dc7f46368982acfa2095ee8d73fe25b5b6e32279900cdb0fd372a3654e41c5", HMAC384Hex("foo", "bar"))
	assert.Equal(t, "HZBw0Hy3dG4GZMzMbOwfqZbcf0Y2iYKs+iCV7o1z/iW1tuMieZAM2w/TcqNlTkHF", HMAC384Base64StdEnc("foo", "bar"))
	assert.Equal(t, "HZBw0Hy3dG4GZMzMbOwfqZbcf0Y2iYKs+iCV7o1z/iW1tuMieZAM2w/TcqNlTkHF", HMAC384Base64RawStdEnc("foo", "bar"))
	assert.Equal(t, "HZBw0Hy3dG4GZMzMbOwfqZbcf0Y2iYKs-iCV7o1z_iW1tuMieZAM2w_TcqNlTkHF", HMAC384Base64URLEnc("foo", "bar"))
	assert.Equal(t, "HZBw0Hy3dG4GZMzMbOwfqZbcf0Y2iYKs-iCV7o1z_iW1tuMieZAM2w_TcqNlTkHF", HMAC384Base64RawURLEnc("foo", "bar"))
}
