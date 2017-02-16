package main

import "fmt"

func main() {
	// BEGIN OMIT
	animals := []string{"fish", "dog"}

	animal := animals[1] // HL
	fmt.Println(animal)

	animals[1] = "cat" // HL
	fmt.Println(animals)
	// END OMIT
}
