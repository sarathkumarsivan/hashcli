package cmd

import (
	"flag"
	"os"

	"github.com/sarathkumarsivan/hashutils/hash"
)

func parseCommandLine() Options {
	flags := flag.NewFlagSet("hash", flag.ExitOnError)
	algorithm := flags.String("algorithm", "sha1", "--algorithm <your text here> default: sha1")
	encoding := flags.String("encoding", "hex", "--encoding <your text here> default: hex")
	text := flags.String("text", "", "--text <your text here>")
	file := flags.String("file", "", "--file <your text here>")
	var options = Options{}
	flags.Parse(os.Args[1:])
	if flags.Parsed() {
		options.algorithm = hash.Algorithm(*algorithm)
		options.encoding = hash.Encoding(*encoding)
		if *text != "" {
			options.text = *text
			options.valid = true
		}
		if *file != "" {
			options.file = *file
			options.valid = true
		}
	}
	if !options.valid {
		Exit("hashcli: not enough options to perform hashing", flags)
	}
	return options
}
