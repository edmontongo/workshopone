package main

import (
	"fmt"
	"math"
)

// BEGIN OMIT
func pedanticRoot(a float64) (float64, float64) {
	root := math.Sqrt(a)
	return root, -root
}

func main() {
	c := 16.0
	r1, r2 := pedanticRoot(c)
	fmt.Println("The roots of", c, "are", r1, "and", r2)
}

// END OMIT
