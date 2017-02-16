package main

import (
	"fmt"
)

func main() {
	// START OMIT
	type Point struct {X, Y int}
	type Circle struct {
		Point
		Radius int
	}
	c := Circle{Point{1, 2}, 3}
	// c = Circle{X: 1, Y: 2, Radius: 3} // Doesn't work.
	fmt.Println(c.X, c.Y, c.Point.X)
	// END OMIT
}
