package main

import "fmt"

func main() {
	// BEGIN OMIT
	fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}

	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune", // HL
	}
	// END OMIT
	fmt.Println(fibonacci)

	fmt.Println(planets)
}
