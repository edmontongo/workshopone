package main

import "fmt"

// BEGIN OMIT
// minMax gives you the low and high values.
func minMax(numbers []int) (min, max int) {
	if len(numbers) < 1 {
		return 0, 0
	}

	min, max = numbers[0], numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func main() {
	fmt.Println(minMax([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))       // 0, 21
	fmt.Println(minMax([]int{18, 44, 45, 97, 81, 2, 19, 19, 27})) // 2, 97
	fmt.Println(minMax([]int{-3, 3, 4, -5}))
	fmt.Println(minMax(nil))
}

// END OMIT
