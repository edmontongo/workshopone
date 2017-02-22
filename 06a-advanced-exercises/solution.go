package main

import (
	"fmt"
)

func main() {
	printTable(2, 7)
}

// BEGIN OMIT
func printTable(a, b int) {
	fmt.Printf("   X|") // Print header...
	for j := a; j <= b; j++ {
		fmt.Printf("%4d", j)
	}
	fmt.Printf("\n")

	fmt.Printf("----+") // Print divider...
	for j := a; j <= b; j++ {
		fmt.Printf("----")
	}
	fmt.Printf("\n")

	for i := a; i <= b; i++ { // Print numbers...
		fmt.Printf("%4d|", i)
		for j := a; j <= b; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Printf("\n")
	}
}

// END OMIT
