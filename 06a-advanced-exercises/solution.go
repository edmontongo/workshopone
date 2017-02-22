package main

import (
	"fmt"
	"strconv"
)

func main() {
	multiplicationTable(2, 7)
}

// BEGIN OMIT
func multiplicationTable(low, high int) {
	maxwidth := len(strconv.Itoa(high*high)) + 1
	highwidth := len(strconv.Itoa(high))
	fmt.Printf("%*s |", highwidth, "X")
	for i := low; i <= high; i++ {
		fmt.Printf("%*d", maxwidth, i)
	}
	fmt.Printf("\n")

	// print lines
	for i := 0; i < highwidth+(maxwidth*(high-low+1))+2; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

	for i := low; i <= high; i++ {
		fmt.Printf("%*d |", highwidth, i)
		for j := low; j <= high; j++ {
			fmt.Printf("%*d", maxwidth, i*j)
		}
		fmt.Printf("\n")
	}
}

// END OMIT
