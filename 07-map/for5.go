package main

import "fmt"

func main() {
	ages := map[string]int{
		"Baby":   0,
		"Justin": 45,
		"Donald": 70,
		"Canada": 150,
	}
	// START OMIT
	var sum int
	var n int
	for _, value := range ages {
		n++
		sum += value
	}
	fmt.Println("Average age rounded down is:", sum/n)
	// END OMIT
}
