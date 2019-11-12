package output

import "fmt"

func PrintShort(info string) {
	fmt.Println("Company info, short version")
	fmt.Println(info)
}

func PrintSummary(info string) {
	fmt.Println("Allabolag.se Summary")
	fmt.Println("====================")
	fmt.Printf("Link: %s\n", info)
}
