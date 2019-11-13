package scrape

type CompanyInfoScraper interface {
	Search(term string) ([]Company, error)
	Details(c Company) (CompanyDetails, error)
}

type Company struct {
	Name string
	Link string
}

type CompanyDetails struct {
	Company
	Fiscal
}

type Fiscal struct {
	Year    string
	Revenue string
}
