package main

import (
	"fmt"
)

func main() {
	// START OMIT
	type Circle struct {
		X, Y, Radius int
	}
	origin := Circle{Radius: 1}
	zero := Circle{}
	fmt.Printf("origin=%v\nzero=%v\n", origin, zero)
	// zeroRadius := Circle{1, 5} // "too few values"
	// END OMIT
}
