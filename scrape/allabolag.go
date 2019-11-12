package scrape

type AllaBolagScraper struct{}

func (s *AllaBolagScraper) Search(term string) ([]string, error) {
	return []string{"https://www.allabolag.se/5566619531/caspeco-ab"}, nil
}

func NewAllaBolagScraper() *AllaBolagScraper {
	return &AllaBolagScraper{}
}
