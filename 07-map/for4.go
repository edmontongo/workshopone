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
	for key, value := range ages {
		fmt.Print(key, ":", value, ", ")
	}
	// END OMIT
}
