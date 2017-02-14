package main

import (
	"fmt"
	"bytes"
)

var _ bytes.Buffer

func main() { 
	// START OMIT
	type Circle struct {X, Y, Radius int}
	small := Circle{X: 1, Y: 1, Radius: 1}
	fmt.Printf("%#v\n", small)

	large := &Circle{2, 2, 2}
	fmt.Printf("%#v\n", large)

	type CircleGroup struct {
		name string
		circles []Circle
	}
	group := CircleGroup{
		"eyes",
		[]Circle{{2, 2, 2}, {5, 2, 2}},
	}
	fmt.Printf("%v\n", group)
	// END OMIT
}
