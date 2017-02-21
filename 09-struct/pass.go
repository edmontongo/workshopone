package main

import (
	"fmt"
)

// START OMIT
type Circle struct {
	X, Y, Radius int
}
func translateCircle(c *Circle, dx, dy int) {
	c.X += dx
	c.Y += dy
	fmt.Printf("translateCircle: %v\n", c)
}
func main() {
	small := &Circle{10, 10, 2}
	translateCircle(small, 3, -3)
	fmt.Printf("main: X=%d Y=%d\n", small.X, small.Y)
}
// END OMIT
