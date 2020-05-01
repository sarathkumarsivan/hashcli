package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarathkumarsivan/hashcli"
)

const (
	flagText      = "text"
	flagTextEmpty = ""
	flagTextDesc  = "Text or string to Hash"
)

func parseCommandLine() map[string]string {
	options := make(map[string]string)
	md5Hash := flag.NewFlagSet(hashcli.AlgorithmMD5, flag.ExitOnError)
	md5HashText := md5Hash.String(flagText, flagTextEmpty, flagTextEmpty)
	sha1Hash := flag.NewFlagSet(hashcli.AlgorithmSHA1, flag.ExitOnError)
	sha1HashText := sha1Hash.String(flagText, flagTextEmpty, flagTextDesc)
	sha256Hash := flag.NewFlagSet(hashcli.AlgorithmSHA256, flag.ExitOnError)
	sha256HashText := sha256Hash.String(flagText, flagTextEmpty, flagTextDesc)
	sha512Hash := flag.NewFlagSet(hashcli.AlgorithmSHA512, flag.ExitOnError)
	sha512HashText := sha512Hash.String(flagText, flagTextEmpty, flagTextDesc)

	if len(os.Args) < 2 {
		fmt.Println("unknown command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case hashcli.AlgorithmMD5:
		md5Hash.Parse(os.Args[2:])
	case hashcli.AlgorithmSHA1:
		sha1Hash.Parse(os.Args[2:])
	case hashcli.AlgorithmSHA256:
		sha256Hash.Parse(os.Args[2:])
	case hashcli.AlgorithmSHA512:
		sha512Hash.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if md5Hash.Parsed() {
		if *md5HashText == flagTextEmpty {
			md5Hash.PrintDefaults()
			os.Exit(1)
		}
		options[hashcli.Algorithm] = hashcli.AlgorithmMD5
		options[flagText] = *md5HashText
	}

	if sha1Hash.Parsed() {
		if *sha1HashText == flagTextEmpty {
			sha1Hash.PrintDefaults()
			os.Exit(1)
		}
		options[hashcli.Algorithm] = hashcli.AlgorithmSHA1
		options[flagText] = *sha1HashText
	}

	if sha256Hash.Parsed() {
		if *sha256HashText == flagTextEmpty {
			sha256Hash.PrintDefaults()
			os.Exit(1)
		}
		options[hashcli.Algorithm] = hashcli.AlgorithmSHA256
		options[flagText] = *sha256HashText
	}

	if sha512Hash.Parsed() {
		if *sha512HashText == flagTextEmpty {
			sha512Hash.PrintDefaults()
			os.Exit(1)
		}
		options[hashcli.Algorithm] = hashcli.AlgorithmSHA512
		options[flagText] = *sha512HashText
	}

	return options
}

func run() {
	options := parseCommandLine()
	hash, err := hashcli.GetHash(options[flagText], options[hashcli.Algorithm])
	if err != nil {
		fmt.Errorf("error hasing text: %s using algorithm %s, error: %s",
			options[hashcli.Algorithm], options[flagText], err)
		os.Exit(1)
	}
	fmt.Printf("%s (%s): %s\n", options[hashcli.Algorithm], options[flagText], hash)
	return
}

func main() {
	run()
}
