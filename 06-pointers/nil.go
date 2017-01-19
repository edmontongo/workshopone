package main

import "fmt"

func main() {
	// BEGIN OMIT
	var a *int
	var b int = 7

	fmt.Printf("Value of a %v\n", a)
	// fmt.Printf("a points to %v\n", *a)

	a = &b
	fmt.Printf("Value of a %v\n", a)
	fmt.Printf("a points to %v\n", *a)

	a = nil
	fmt.Printf("Value of a %v\n", a)
	// fmt.Printf("a points to %v\n", *a)
	// END OMIT
}
