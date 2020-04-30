package hashcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSHA1Hash(t *testing.T) {
	assert.NotEmpty(t, GetSHA1Hash("abc"))
}
