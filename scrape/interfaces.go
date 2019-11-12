package scrape

type CompanyInfoScraper interface {
	Search(term string) ([]Company, error)
}

type Company struct {
	Name string
	Link string
}
