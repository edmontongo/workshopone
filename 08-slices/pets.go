package main

import "fmt"

// BEGIN OMIT
// filterPets lets you know which animals are considered pets.
func filterPets(animals []string) (pets []string) {
	for _, a := range animals {
		switch a {
		case "cat", "fish", "bunny", "hamster", "dog":
			pets = append(pets, a)
		}
	}
	return pets
}

func main() {
	fmt.Println(filterPets([]string{"cat", "chicken"}))         // cat
	fmt.Println(filterPets([]string{"fish", "mouse", "bunny"})) // fish, bunny
	fmt.Println(filterPets([]string{"hamster", "cow", "dog", "moose"}))
}

// END OMIT
