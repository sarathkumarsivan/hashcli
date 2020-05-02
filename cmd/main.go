package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarathkumarsivan/hashcli"
)

const (
	flagText  = "text"
	flagFile  = "file"
	flagEmpty = ""
)

func parseCommandLine() map[string]string {
	options := make(map[string]string)

	md5Hash := flag.NewFlagSet(hashcli.MD5Hash, flag.ExitOnError)
	md5HashText := md5Hash.String(flagText, flagEmpty, "--text <your text here>")
	md5HashFile := md5Hash.String(flagFile, flagEmpty, "--file <your file here>")

	sha1Hash := flag.NewFlagSet(hashcli.SHA1Hash, flag.ExitOnError)
	sha1HashText := sha1Hash.String(flagText, flagEmpty, "--text <your text here>")
	sha1HashFile := sha1Hash.String(flagFile, flagEmpty, "--file <your file here>")

	sha256Hash := flag.NewFlagSet(hashcli.SHA256Hash, flag.ExitOnError)
	sha256HashText := sha256Hash.String(flagText, flagEmpty, "--text <your text here>")
	sha256HashFile := sha256Hash.String(flagFile, flagEmpty, "--file <your file here>")

	sha512Hash := flag.NewFlagSet(hashcli.SHA512Hash, flag.ExitOnError)
	sha512HashText := sha512Hash.String(flagText, flagEmpty, "--text <your text here>")
	sha512HashFile := sha512Hash.String(flagFile, flagEmpty, "--file <your file here>")

	if len(os.Args) < 2 {
		fmt.Println("unknown command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case hashcli.MD5Hash:
		md5Hash.Parse(os.Args[2:])
	case hashcli.SHA1Hash:
		sha1Hash.Parse(os.Args[2:])
	case hashcli.SHA256Hash:
		sha256Hash.Parse(os.Args[2:])
	case hashcli.SHA512Hash:
		sha512Hash.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if md5Hash.Parsed() {
		if *md5HashText != flagEmpty {
			options[hashcli.Algorithm] = hashcli.MD5Hash
			options[flagText] = *md5HashText
		} else if *md5HashFile != flagEmpty {
			options[hashcli.Algorithm] = hashcli.MD5Hash
			options[flagFile] = *md5HashFile
		} else {
			md5Hash.PrintDefaults()
			os.Exit(1)
		}
	}

	if sha1Hash.Parsed() {
		if *sha1HashText != flagEmpty {
			options[hashcli.Algorithm] = hashcli.SHA1Hash
			options[flagText] = *sha1HashText
		} else if *sha1HashFile != flagEmpty {
			options[hashcli.Algorithm] = hashcli.SHA1Hash
			options[flagFile] = *sha1HashFile
		} else {
			sha1Hash.PrintDefaults()
			os.Exit(1)
		}
	}

	if sha256Hash.Parsed() {
		if *sha256HashText != flagEmpty {
			options[hashcli.Algorithm] = hashcli.SHA256Hash
			options[flagText] = *sha256HashText
		} else if *sha256HashFile != flagEmpty {
			options[hashcli.Algorithm] = hashcli.SHA256Hash
			options[flagFile] = *sha256HashFile
		} else {
			sha256Hash.PrintDefaults()
			os.Exit(1)
		}
	}

	if sha512Hash.Parsed() {
		if *sha512HashText != flagEmpty {
			options[hashcli.Algorithm] = hashcli.SHA512Hash
			options[flagText] = *sha512HashText
		} else if *sha512HashFile != flagEmpty {
			options[hashcli.Algorithm] = hashcli.SHA512Hash
			options[flagFile] = *sha512HashFile
		} else {
			sha512Hash.PrintDefaults()
			os.Exit(1)
		}
	}

	return options
}

func run() {
	options := parseCommandLine()
	if options[flagText] != flagEmpty {
		hash, err := hashcli.MakeHash(options[flagText], options[hashcli.Algorithm])
		if err != nil {
			fmt.Errorf("error hasing text: %s using algorithm %s, error: %s",
				options[hashcli.Algorithm], options[flagText], err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options[hashcli.Algorithm], options[flagText], hash)
	}
	if options[flagFile] != flagEmpty {
		hash, err := hashcli.MakeFileHash(options[flagFile], options[hashcli.Algorithm])
		if err != nil {
			fmt.Errorf("error hasing file: %s using algorithm %s, error: %s",
				options[hashcli.Algorithm], options[flagFile], err)
			os.Exit(1)
		}
		fmt.Printf("%s (%s): %s\n", options[hashcli.Algorithm], options[flagFile], hash)
	}

	return
}

func main() {
	run()
}
