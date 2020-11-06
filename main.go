package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var format string

func main() {
	flag.StringVar(&format, "format", "yaml", "the changelog format. can be yaml or json")
	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Println("missing changelog filepath")
		os.Exit(1)
	}

	changelogFilepath := os.Args[1]
	changelogRaw, err := ioutil.ReadFile(changelogFilepath)
	if err != nil {
		fmt.Printf("cannot read file %q\n", changelogFilepath)
		os.Exit(1)
	}

	diagnostics := Validate(changelogRaw)
	if diagnostics.HasErrors() {
		diagnostics.Print(os.Stdout)
		os.Exit(1)
	}
}
