package output

import (
	"fmt"

	"github.com/vertan/allabolag-cli/scrape"
)

func PrintShort(info string) {
	fmt.Println("Company info, short version")
	fmt.Println(info)
}

func PrintSummary(c scrape.CompanyDetails) {
	fmt.Println("Allabolag.se Summary")
	fmt.Println("====================")
	fmt.Printf("Name: %s\n", c.Company.Name)
	fmt.Printf("Revenue: %s K SEK\n", c.Fiscal.Revenue)
	fmt.Printf("Link: %s\n", c.Company.Link)
}
