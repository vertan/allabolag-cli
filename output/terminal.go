package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vertan/allabolag-cli/scrape"
)

func PrintTerse(c scrape.CompanyDetails) {
	fmt.Printf("%s\n", c.Company.Name)
	fmt.Printf("%s\n", c.Company.Link)
	fmt.Printf("Revenue (%d): %dk\n", c.FiscalDetails[0].Year, c.FiscalDetails[0].Revenue)
	fmt.Printf("Results (%d): %dk\n", c.FiscalDetails[0].Year, c.FiscalDetails[0].Result)
}

func PrintSummary(c scrape.CompanyDetails) {
	fmt.Printf("%s\n", c.Company.Name)
	fmt.Printf("%s\n", c.Company.Link)
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
