package main

import "fmt"

func main() {
	// BEGIN OMIT
	numbers := []string{"zero", "one", "two", "three"}

	for i, word := range numbers { // HL
		fmt.Println(i, word)
	}
	// END OMIT
}
