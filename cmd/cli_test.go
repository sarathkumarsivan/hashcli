package cmd

import (
	"flag"
	"testing"

	"github.com/sarathkumarsivan/hashutils/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCommandLine(t *testing.T) {
	args := []string{"hash", "-t", "foo"}
	options, err := ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "md5", "-t", "foo", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("md5"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha1", "-t", "foo", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha224", "-t", "foo", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha224"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha256", "-t", "foo", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha256"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha384", "-t", "foo", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha384"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha512", "-t", "foo", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha512"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "md5", "-t", "foo", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("md5"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha1", "-t", "foo", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha224", "-t", "foo", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha224"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha256", "-t", "foo", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha256"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha384", "-t", "foo", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha384"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-a", "sha512", "-t", "foo", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha512"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo", options.text)

	args = []string{"hash", "-f", "foo.txt"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha1", "-f", "foo.txt", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha224", "-f", "foo.txt", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha224"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha256", "-f", "foo.txt", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha256"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha384", "-f", "foo.txt", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha384"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha512", "-f", "foo.txt", "-e", "hex"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha512"), options.algorithm)
	assert.Equal(t, hash.Encoding("hex"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha1", "-f", "foo.txt", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha1"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha224", "-f", "foo.txt", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha224"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha256", "-f", "foo.txt", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha256"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha384", "-f", "foo.txt", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha384"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)

	args = []string{"hash", "-a", "sha512", "-f", "foo.txt", "-e", "base64"}
	options, err = ParseCommandLine(args, flag.ContinueOnError)
	require.NoError(t, err, "Error parsing commandline options")
	assert.Equal(t, hash.Algorithm("sha512"), options.algorithm)
	assert.Equal(t, hash.Encoding("base64"), options.encoding)
	assert.Equal(t, "foo.txt", options.file)
}
