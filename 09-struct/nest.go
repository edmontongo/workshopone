package main

import (
	"fmt"
)

func main() {
	// START OMIT
	type Point struct {X, Y int}
	type Circle struct {
		Coords Point
		Radius int
	}
	c := Circle{Point{1, 2}, 3}
	fmt.Println(c.Coords.X, c.Coords.Y)
	// END OMIT
}
