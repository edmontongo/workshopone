package main

import (
	"fmt"
)

// START STRUCT OMIT
type Circle struct {
	X, Y, Radius int
}
// END STRUCT OMIT

func main() { 
// START BASIC OMIT
	feet := 5.68
	meters := 1.73
	fmt.Println(feet + meters) // Is the result meaningful?

	type Feet float64
	type Meters float64
	mile := Feet(5280)
	km := Meters(1000)
	// fmt.Println(mile + km) // type mismatch
// END BASIC OMIT
	_ = mile
	_ = km
} 
