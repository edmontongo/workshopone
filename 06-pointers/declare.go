package main

import "fmt"

func main() {
	var a int = 5
	var b *int = &a
	c := &a

	fmt.Printf("Address of a %v, value %v\n", &a, a)
	fmt.Printf("Address of b %v, value %v\n", &b, b)
	fmt.Printf("Address of c %v, value %v\n", &c, c)
}
