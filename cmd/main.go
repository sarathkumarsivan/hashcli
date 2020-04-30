package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarathkumarsivan/hashcli"
)

const (
	HashingAlgorithm = "algorithm"
	AlgorithmSHA1    = "sha1"
	AlgorithmMD5     = "md5"
	FlagText         = "text"
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
	case AlgorithmSHA1:
		sha1Hash.Parse(os.Args[2:])
	case AlgorithmMD5:
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
		options[HashingAlgorithm] = AlgorithmSHA1
		options["text"] = *sha1HashText
	}

	return options
}

func main() {
	options := parseCommandLine()
	if options[HashingAlgorithm] == AlgorithmSHA1 {
		hash := hashcli.GetSHA1Hash(options[FlagText])
		fmt.Printf("sha-1 (%s): %x\n", options[FlagText], hash)
		return
	}
	if options[HashingAlgorithm] == AlgorithmSHA1 {
		hash := hashcli.GetMD5Hash(options[FlagText])
		fmt.Printf("md5 (%s): %x\n", options[FlagText], hash)
		return
	}
}
