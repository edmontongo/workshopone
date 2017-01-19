package main

import "fmt"

// BEGIN OMIT
func main() {
	var c int = 5
	mutate(&c)
	fmt.Println("Mutate")
}

func mutate(b *int) {
	*b += 5
}

func copy(b int) {
	*b += 5
}

// END OMIT
