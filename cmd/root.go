package cmd

import (
	"encoding/json"
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
	pretty    bool
}

type response struct {
	Text      string `json:"text,omitempty"`
	File      string `json:"file,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
	Encoding  string `json:"encoding,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

func printAsJSON(response response) {
	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func printAsPrettyJSON(response response) {
	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func Exit(message string, flags *flag.FlagSet) {
	fmt.Println(message)
	flags.PrintDefaults()
	os.Exit(1)
}

func Execute() {
	options, err := ParseCommandLine(os.Args, flag.ExitOnError)
	if err != nil {
		os.Exit(1)
	}
	response := response{
		Text:      options.text,
		File:      options.file,
		Algorithm: string(options.algorithm),
		Encoding:  string(options.encoding),
	}

	if options.text != "" {
		maker := hash.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := maker.HashText(options.text)
		if err != nil {
			fmt.Errorf("error hasing text: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		response.Hash = hash
	}
	if options.file != "" {
		maker := hash.New().Algorithm(options.algorithm).Encoding(options.encoding).Build()
		hash, err := maker.HashFile(options.file)
		if err != nil {
			fmt.Errorf("error hasing file: %s using algorithm %s, error: %s", options.text, options.algorithm, err)
			os.Exit(1)
		}
		response.Hash = hash
	}

	if options.pretty {
		printAsPrettyJSON(response)
	} else {
		printAsJSON(response)
	}
}
