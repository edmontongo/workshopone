package main

import "fmt"

func main() {
	// BEGIN OMIT
	pets := []string{"cat", "fish"}

	pets = append(pets, "dog") // HL

	pets = append(pets, "bunny", "hamster") // HL

	fmt.Println(pets)
	// END OMIT
}
