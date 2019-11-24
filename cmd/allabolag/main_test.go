package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vertan/allabolag-cli/scrape"
)

// TODO: Move out to mock package?
type MockScraper struct {
	SearchFunc  func(string) ([]scrape.Company, error)
	DetailsFunc func(scrape.Company) (*scrape.CompanyDetails, error)
}

func (s MockScraper) Search(term string) ([]scrape.Company, error) { return s.SearchFunc(term) }
func (s MockScraper) Details(c scrape.Company) (*scrape.CompanyDetails, error) {
	return s.DetailsFunc(c)
}

func TestNoCompanyResultsReturnsNoError(t *testing.T) {
	noResultSearchFunc := func(term string) ([]scrape.Company, error) { return []scrape.Company{}, nil }
	emptyDetailsFunc := func(c scrape.Company) (*scrape.CompanyDetails, error) { return nil, nil }
	s := MockScraper{SearchFunc: noResultSearchFunc, DetailsFunc: emptyDetailsFunc}

	runFunc := func() { run(s, "anything with no results", false) }
	assert.NotPanics(t, runFunc)
}
