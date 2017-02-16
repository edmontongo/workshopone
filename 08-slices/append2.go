package main

import "fmt"

func main() {
	// BEGIN OMIT
	pets := []string{"cat", "fish", "dog"}

	animals := append(pets, "fox", "tiger") // HL

	fmt.Println(animals)
	fmt.Println(pets)
	// END OMIT
}
