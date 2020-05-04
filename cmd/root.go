package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarathkumarsivan/hashutils/hash"
)

type Options struct {
	algorithm hash.Algorithm
	encoding  hash.Encoding
	text      string
	file      string
	valid     bool
}

func Exit(message string, flags *flag.FlagSet) {
	fmt.Println(message)
	flags.PrintDefaults()
	os.Exit(1)
}

func Run() {
	options := parseCommandLine()
	if options.text != "" {
		hashcli := hash.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := hashcli.HashText(options.text)
		if err != nil {
			fmt.Errorf("error hasing text: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options.algorithm, options.text, hash)
	}
	if options.file != "" {
		hashcli := hash.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := hashcli.HashFile(options.file)
		if err != nil {
			fmt.Errorf("error hasing file: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options.algorithm, options.file, hash)
	}
}
