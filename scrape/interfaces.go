package scrape

type CompanyInfoScraper interface {
	Search(term string) ([]Company, error)
	Details(c Company) (*CompanyDetails, error)
}

type Company struct {
	Name string
	Link string
}

type CompanyDetails struct {
	Company
	FiscalDetails []FiscalDetails
}

type FiscalDetails struct {
	Year    int
	Revenue int
	Result  int
}
