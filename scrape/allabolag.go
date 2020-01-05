package scrape

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

// AllaBolagScraper is an implementation of CompanyInfoScraper fetching info from allabolag.se.
type AllaBolagScraper struct{}

const (
	searchURL       = "https://www.allabolag.se/what/%s"
	maxYearsToFetch = 5
)

type ParsedCompany struct {
	Name string `json:"jurnamn"`
	Link string `json:"linkTo"`
}

// Search takes a search term as a parameter and searches allabolag.se for companies.
func (s *AllaBolagScraper) Search(term string) ([]Company, error) {
	c := colly.NewCollector()
	companies := []Company{}

	re := regexp.MustCompile(`(.+/).+$`)

	c.OnHTML("search", func(e *colly.HTMLElement) {
		raw := e.Attr(":search-result-default")
		parsed := []ParsedCompany{}
		json.Unmarshal([]byte(raw), &parsed)

		for _, c := range parsed {
			link := fmt.Sprintf("https://www.allabolag.se/%s", re.FindStringSubmatch(c.Link)[1])
			comp := Company{Name: c.Name, Link: link}
			companies = append(companies, comp)
		}
	})

	_ = c.Visit(fmt.Sprintf(searchURL, term))

	return companies, nil
}

// Details returns details about a specific company.
func (s *AllaBolagScraper) Details(comp Company) (*CompanyDetails, error) {
	c := colly.NewCollector()
	details := CompanyDetails{}
	details.Company = comp

	years := []string{}
	figures := []string{}

	c.OnHTML(".chart__label", func(e *colly.HTMLElement) {
		years = append(years, e.Text)
	})

	c.OnHTML(".chart__data", func(e *colly.HTMLElement) {
		figures = append(figures, e.Attr("value"))
	})
	_ = c.Visit(fmt.Sprintf("%sbokslut", comp.Link))

	yearsToFetch := min(len(years), maxYearsToFetch)
	fiscalDetails, err := convertToFiscalDetails(years[:yearsToFetch], figures[:yearsToFetch*2])
	if err != nil {
		return nil, err
	}

	details.FiscalDetails = fiscalDetails

	return &details, nil
}

// NewAllaBolagScraper returns a new AllaBolagScraper
func NewAllaBolagScraper() *AllaBolagScraper {
	return &AllaBolagScraper{}
}

func convertToFiscalDetails(years, figures []string) ([]FiscalDetails, error) {
	converted := []FiscalDetails{}

	pad := 0
	for _, v := range years {
		y, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("year conversion failed")
		}
		rev, err := strconv.Atoi(figures[pad])
		if err != nil {
			return nil, errors.New("revenue conversion failed")
		}
		res, err := strconv.Atoi(figures[pad+1])
		if err != nil {
			return nil, errors.New("result conversion failed")
		}

		fd := FiscalDetails{
			Year:    y,
			Revenue: rev,
			Result:  res,
		}
		converted = append(converted, fd)
		pad += 2
	}

	return converted, nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
