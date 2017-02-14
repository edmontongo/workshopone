package main

import (
	"fmt"
)

func main() { 
	// START OMIT
	// Comparable
	type Circle struct {
		X, Y, Radius int
	}
	small := Circle{X: 10, Y: 10, Radius: 2}
	big := Circle{10, 10, 5}
	fmt.Println("small == big: ", small == big)

	// Not comparable
	type Group struct {
		name string
		members []string
	}
	fruits := Group{name: "fruits", members: []string{"apple", "pear"}}
	rocks := Group{name: "rocks", members: []string{"basalt", "granite"}}
	// fmt.Println("fruits == rocks: ", fruits == rocks) // invalid operation
	// END OMIT
	_ = fruits
	_ = rocks
}
