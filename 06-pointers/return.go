package main

import "fmt"

// BEGIN OMIT
func main() {
	var c *int = allocate()
	fmt.Printf("c has address %v and value %v and points to %v\n", &c, c, *c)
}

func allocate() *int {
	a := 1
	b := 2
	fmt.Printf("a has address %v and value %v\n", &a, a)
	fmt.Printf("b has address %v and value %v\n", &b, b)

	return &b
}

// END OMIT
