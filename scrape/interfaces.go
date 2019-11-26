package scrape

// CompanyInfoScraper represents a generic scraper of company information.
type CompanyInfoScraper interface {
	Search(term string) ([]Company, error)
	Details(c Company) (*CompanyDetails, error)
}

// Company represents a company.
type Company struct {
	Name string
	Link string
}

// CompanyDetails represents further details about a company.
type CompanyDetails struct {
	Company
	FiscalDetails []FiscalDetails
}

// FiscalDetails represents the financial information about a company for one year.
type FiscalDetails struct {
	Year    int
	Revenue int
	Result  int
}
