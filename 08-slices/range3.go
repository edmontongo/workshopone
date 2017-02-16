package main

import "fmt"

func main() {
	// BEGIN OMIT
	numbers := []string{"zero", "one", "two", "three"}

	for _, n := range numbers { // HL
		fmt.Println(n)
	}
	// END OMIT
}
