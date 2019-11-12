package main

import (
	"flag"

	"github.com/vertan/allabolag-cli/output"
)

const minPositionalArgs = 1

func main() {
	// Parse flags
	shortMode := flag.Bool("t", false, "print company information in terse form")
	flag.Parse()

	// Company search term is a required argument
	if flag.NArg() < minPositionalArgs {
		flag.Usage()
	}

	searchTerm := flag.Arg(0)

	if *shortMode {
		output.PrintShort(searchTerm)
	}

}
