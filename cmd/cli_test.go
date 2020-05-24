package cmd

import (
	"flag"
	"testing"

	"github.com/sarathkumarsivan/hashutils/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCommandLine(t *testing.T) {
	args := []string{"hash", "-a", "sha1", "-t", "foo", "-e", "hex"}
	options, err := ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha1", "-f", "foo.txt", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)
}
