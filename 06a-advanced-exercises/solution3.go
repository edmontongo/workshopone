package main

import "fmt"

// BEGIN OMIT
func calcPi(iterations int) float64 {
	// This function won't do the exact number of iterations
	// If iterations is an even number, iterations-1 is used
	pi := float64(0)
	for i := 1; i < iterations/2; {
		pi += 1 / float64(i)
		i += 2
		pi -= 1 / float64(i)
		i += 2
	}

	return 4 * pi
}

func main() {
	fmt.Println(calcPi(100))
	fmt.Println(calcPi(1000))
	fmt.Println(calcPi(10000))
}

// END OMIT
