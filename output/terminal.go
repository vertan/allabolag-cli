package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vertan/allabolag-cli/scrape"
)

func PrintShort(info string) {
	fmt.Println("Company info, short version")
	fmt.Println(info)
}

func PrintSummary(c scrape.CompanyDetails) {
	fmt.Printf("Name: %s\n", c.Company.Name)
	fmt.Printf("Link: %s\n", c.Company.Link)
	fmt.Println("--------------------")
	printFiscalTable(c.FiscalDetails)
}

func printFiscalTable(details []scrape.FiscalDetails) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Year\tRevenue\tResult")
	for _, v := range details {
		fmt.Fprintln(w, fmt.Sprintf("%d\t%dk\t%dk", v.Year, v.Revenue, v.Result))
	}
	w.Flush()
}
