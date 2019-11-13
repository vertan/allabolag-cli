package scrape

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type AllaBolagScraper struct{}

const searchUrl = "https://www.allabolag.se/what/%s"

func (s *AllaBolagScraper) Search(term string) ([]Company, error) {
	c := colly.NewCollector()
	companies := []Company{}

	c.OnHTML(".search-results__item__title a[href]", func(e *colly.HTMLElement) {
		name := e.Text
		link := e.Attr("href")
		comp := Company{Name: name, Link: link}
		companies = append(companies, comp)
	})

	c.Visit(fmt.Sprintf(searchUrl, term))

	return companies, nil
}

func (s *AllaBolagScraper) Details(comp Company) (CompanyDetails, error) {
	c := colly.NewCollector()
	details := CompanyDetails{}
	details.Company = comp

	nums := []string{}

	c.OnHTML(".company-account-figures td", func(e *colly.HTMLElement) {
		figure := strings.Trim(e.Text, "\t \n")
		nums = append(nums, figure)
	})
	c.Visit(comp.Link)

	details.Fiscal.Revenue = nums[0]

	return details, nil
}

func NewAllaBolagScraper() *AllaBolagScraper {
	return &AllaBolagScraper{}
}
