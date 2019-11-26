package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vertan/allabolag-cli/scrape"
)

// PrintTerse outputs company details in a terse format to the terminal.
func PrintTerse(c scrape.CompanyDetails) {
	fmt.Printf("%s\n", c.Company.Name)
	fmt.Printf("%s\n", c.Company.Link)

	if len(c.FiscalDetails) > 0 {
		fmt.Printf("Revenue (%d): %dk\n", c.FiscalDetails[0].Year, c.FiscalDetails[0].Revenue)
		fmt.Printf("Results (%d): %dk\n", c.FiscalDetails[0].Year, c.FiscalDetails[0].Result)
	}
}

// PrintSummary outputs company details in a summary format to the terminal.
func PrintSummary(c scrape.CompanyDetails) {
	fmt.Printf("%s\n", c.Company.Name)
	fmt.Printf("%s\n", c.Company.Link)
	if len(c.FiscalDetails) > 0 {
		fmt.Println("--------------------")
		printFiscalTable(c.FiscalDetails)
	}
}

// PrintNoResult outputs a string for when there's no results..
func PrintNoResult(t string) {
	fmt.Printf("No result found for search term %s\n", t)
}

// printFiscalTable renders tabular financial data to the terminal.
func printFiscalTable(details []scrape.FiscalDetails) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Year\tRevenue\tResult")
	for _, v := range details {
		fmt.Fprintln(w, fmt.Sprintf("%d\t%dk\t%dk", v.Year, v.Revenue, v.Result))
	}
	w.Flush()
}
