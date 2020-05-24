package cmd

import (
	"flag"
	"os"

	"github.com/sarathkumarsivan/hashutils/hash"
	"github.com/sarathkumarsivan/hashutils/util"
)

const (
	FlagDescAlgorithm = "Algorithm to be used to hash your text/file/directory."
	FlagDescEncoding  = "Encoding to be used to encode the checksum."
	FlagDescText      = "Text to be hashed with the specified algorithm and encoding."
	FlagDescFile      = "File to be hashed with the specified algorithm and encoding."

	ErrMsgNotEnoughOptions = "hashutils: not enough options to perform hashing"
)

func parseOptions() Options {
	flags := flag.NewFlagSet("hash", flag.ExitOnError)
	algorithm := flags.String("a", "sha1", FlagDescAlgorithm)
	encoding := flags.String("e", "hex", FlagDescEncoding)
	text := flags.String("t", "", FlagDescText)
	file := flags.String("f", "", FlagDescFile)

	var options = Options{}
	flags.Parse(os.Args[1:])

	if flags.Parsed() {
		options.algorithm = hash.Algorithm(*algorithm)
		options.encoding = hash.Encoding(*encoding)
		if util.IsNotEmpty(*text) {
			options.text = *text
			options.valid = true
		}
		if util.IsNotEmpty(*file) {
			options.file = *file
			options.valid = true
		}
	}
	if !options.valid {
		Exit(ErrMsgNotEnoughOptions, flags)
	}
	return options
}
