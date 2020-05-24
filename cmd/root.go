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

func Execute() {
	options := parseOptions()
	if options.text != "" {
		maker := hash.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := maker.HashText(options.text)
		if err != nil {
			fmt.Errorf("error hasing text: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options.algorithm, options.text, hash)
	}
	if options.file != "" {
		maker := hash.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := maker.HashFile(options.file)
		if err != nil {
			fmt.Errorf("error hasing file: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options.algorithm, options.file, hash)
	}
}
