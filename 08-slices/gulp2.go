package main

import "fmt"

func main() {
	// BEGIN OMIT
	animals := []string{"cat", "chicken", "fish", "mouse"}

	hungryCat := animals[:2] // HL
	fmt.Println(hungryCat)

	rous := animals[3:] // HL
	fmt.Println(rous)

	fmt.Println(animals)
	// END OMIT
}
