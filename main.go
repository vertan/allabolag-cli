package main

import (
	"flag"
	"fmt"
)

const minPositionalArgs = 1

func main() {
	fmt.Println("allabolag-cli")

	shortMode := flag.Bool("t", false, "print company information in terse form")
	flag.Parse()

	if flag.NArg() < minPositionalArgs {
		flag.Usage()
	}

	if *shortMode {
		output.printShort("Caspeco AB")
	}

}
