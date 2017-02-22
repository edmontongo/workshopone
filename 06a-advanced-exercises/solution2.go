package main

import "fmt"

func main() {
	for i := uint(1); i <= 7; i++ {
		fmt.Println(getFib(i))
	}
}

// BEGIN OMIT
func getFib(f uint) int {
	if f <= 2 {
		return 1
	}

	x, y := 1, 1 // First two values
	f -= 2       // reduce number of iterations since first two are 1

	for ; f > 0; f-- {
		x, y = y, x+y
	}
	return y
}

// END OMIT
