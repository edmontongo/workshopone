package main

import "fmt"

// START OMIT
func getIncomeCategory(income float64) string {
	switch {
	case income <= 20000:
		return "too damn poor"
	case income <= 40000:
		return "quite poor"
	case income <= 60000:
		return "quite livable"
	case income <= 80000:
		return "just about right"
	case income <= 100000:
		return "starting to cause problems"
	default:
		return "too much money"
	}
}

func main() {
	income := 50000.00
	fmt.Println(income, " is ", getIncomeCategory(income))
}

// END OMIT
