package main

import "fmt"

// BEGIN OMIT

func addInt(a, b int) int {
	return a + b
}

func main() {
	c := 5
	sum := addInt(c, 2)
	fmt.Println("The sum is", sum)
}

// END OMIT
