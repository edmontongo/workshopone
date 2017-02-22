package main

import (
	"fmt"
)

func main() {
	printTable(2, 7)
}

// BEGIN OMIT
func printTable(a, b int) {
	// Print header
	fmt.Printf("   X|")
	for j := a; j <= b; j++ {
		fmt.Printf("%4d", j)
	}
	fmt.Printf("\n")

	// Print divider
	fmt.Printf("----+")
	for j := a; j <= b; j++ {
		fmt.Printf("----")
	}
	fmt.Printf("\n")

	// Print numbers
	for i := a; i <= b; i++ {
		fmt.Printf("%4d|", i)
		for j := a; j <= b; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Printf("\n")
	}
}

// END OMIT
