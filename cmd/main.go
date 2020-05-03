package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarathkumarsivan/hashcli"
)

type Options struct {
	algorithm hashcli.Algorithm
	encoding  hashcli.Encoding
	text      string
	file      string
	valid     bool
}

func parseCommandLine() Options {
	flags := flag.NewFlagSet("hash", flag.ExitOnError)
	algorithm := flags.String("algorithm", "sha1", "--algorithm <your text here> default: sha1")
	encoding := flags.String("encoding", "hex", "--encoding <your text here> default: hex")
	text := flags.String("text", "", "--text <your text here>")
	file := flags.String("file", "", "--file <your text here>")
	var options = Options{}
	flags.Parse(os.Args[1:])
	if flags.Parsed() {
		options.algorithm = hashcli.Algorithm(*algorithm)
		options.encoding = hashcli.Encoding(*encoding)
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
		exit("hashcli: not enough options to perform hashing", flags)
	}
	return options
}

func exit(message string, flags *flag.FlagSet) {
	fmt.Println(message)
	flags.PrintDefaults()
	os.Exit(1)
}

func makeHash() {
	options := parseCommandLine()
	if options.text != "" {
		hashcli := hashcli.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := hashcli.HashText(options.text)
		if err != nil {
			fmt.Errorf("error hasing text: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options.algorithm, options.text, hash)
	}
	if options.file != "" {
		hashcli := hashcli.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := hashcli.HashFile(options.file)
		if err != nil {
			fmt.Errorf("error hasing file: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options.algorithm, options.file, hash)
	}
}

func main() {
	makeHash()
}
