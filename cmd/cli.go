package cmd

import (
	"flag"
	"github.com/sarathkumarsivan/hashutils/hash"
	"github.com/sarathkumarsivan/hashutils/util"
)

const (
	FlagDescAlgorithm = "Algorithm to be used to hash your text/file/directory."
	FlagDescEncoding  = "Encoding to be used to encode the checksum."
	FlagDescText      = "Text to be hashed with the specified algorithm and encoding."
	FlagDescFile      = "File to be hashed with the specified algorithm and encoding."
	FlagDescPretty    = "Specify pretty flag if you want formatted JSON."

	ErrMsgNotEnoughOptions = "hashutils: not enough options to perform hashing"
)

func ParseCommandLine(args []string, errorHandling flag.ErrorHandling) (options Options, err error) {
	flags := flag.NewFlagSet(args[0], errorHandling)
	algorithm := flags.String("a", "sha1", FlagDescAlgorithm)
	encoding := flags.String("e", "hex", FlagDescEncoding)
	text := flags.String("t", "", FlagDescText)
	file := flags.String("f", "", FlagDescFile)
	pretty := flags.Bool("p", false, FlagDescPretty)

	if err = flags.Parse(args[1:]); err != nil {
		return
	}

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
		options.pretty = *pretty
	}
	if !options.valid {
		Exit(ErrMsgNotEnoughOptions, flags)
	}
	return
}
