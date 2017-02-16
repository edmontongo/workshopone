package main

import "fmt"

func main() {
	// BEGIN OMIT
	animals := []string{"cat", "chicken", "fish", "mouse"}

	hungryCat := animals[0:2] // HL
	fmt.Println(hungryCat)

	bigFish := animals[2:3] // HL
	fmt.Println(bigFish)
	// END OMIT
}
