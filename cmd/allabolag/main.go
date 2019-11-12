package main

import (
	"flag"
	"log"

	"github.com/vertan/allabolag-cli/output"
	"github.com/vertan/allabolag-cli/scrape"
)

const minPositionalArgs = 1

func main() {
	// Parse flags
	shortMode := flag.Bool("t", false, "print company information in terse form")
	flag.Parse()

	// Company search term is a required argument
	if flag.NArg() < minPositionalArgs {
		flag.Usage()
		log.Fatal("missing required argument: search term")
	}

	searchTerm := flag.Arg(0)

	if *shortMode {
		output.PrintShort(searchTerm)
	}

	scraper := scrape.NewAllaBolagScraper()

	links, _ := scraper.Search(searchTerm)
	output.PrintSummary(links[0])
}
