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
	terseMode := flag.Bool("t", false, "print company information in terse form")
	flag.Parse()

	// Company search term is a required argument
	if flag.NArg() < minPositionalArgs {
		flag.Usage()
		log.Fatal("missing required argument: search term")
	}

	searchTerm := flag.Arg(0)
	scraper := scrape.NewAllaBolagScraper()

	companies, _ := scraper.Search(searchTerm)
	company, _ := scraper.Details(companies[0])
	if *terseMode {
		output.PrintTerse(*company)
	} else {
		output.PrintSummary(*company)
	}
}
