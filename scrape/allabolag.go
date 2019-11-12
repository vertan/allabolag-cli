package scrape

import (
	"fmt"

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

func NewAllaBolagScraper() *AllaBolagScraper {
	return &AllaBolagScraper{}
}
