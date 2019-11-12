package output

import (
	"fmt"

	"github.com/vertan/allabolag-cli/scrape"
)

func PrintShort(info string) {
	fmt.Println("Company info, short version")
	fmt.Println(info)
}

func PrintSummary(c scrape.Company) {
	fmt.Println("Allabolag.se Summary")
	fmt.Println("====================")
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Link: %s\n", c.Link)
}
