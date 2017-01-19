package main

import "fmt"

func main() {
	// BEGIN OMIT
	var a int = 5
	var b *int = &a

	fmt.Printf("a has value %v\n", a)
	fmt.Printf("b points to value %v\n", *b)

	*b = 6
	fmt.Printf("a has value %v\n", a)
	fmt.Printf("b points to value %v\n", *b)
	// END OMIT
}
