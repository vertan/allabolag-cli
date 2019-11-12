package scrape

type CompanyInfoScraper interface {
	Search(term string) ([]string, error)
}
