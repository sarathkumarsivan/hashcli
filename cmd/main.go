package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarathkumarsivan/hashcli"
)

func parseCommandLine() map[string]string {
	options := make(map[string]string)
	sha1Hash := flag.NewFlagSet("sha1", flag.ExitOnError)
	sha1HashText := sha1Hash.String("text", "", "Text to hash")

	if len(os.Args) < 2 {
		fmt.Println("unknown command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "sha1":
		sha1Hash.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if sha1Hash.Parsed() {
		if *sha1HashText == "" {
			sha1Hash.PrintDefaults()
			os.Exit(1)
		}
		options["algorithm"] = "sha1"
		options["text"] = *sha1HashText
	}

	return options
}

func main() {
	options := parseCommandLine()
	if options["algorithm"] == "sha1" {
		hash := hashcli.GetSHA1Hash(options["text"])
		fmt.Printf("sha-1 (%s): %x\n", options["text"], hash)
		return
	}
	if options["algorithm"] == "md5" {
		hash := hashcli.GetMD5Hash(options["text"])
		fmt.Printf("md5 (%s): %x\n", options["text"], hash)
		return
	}
}
