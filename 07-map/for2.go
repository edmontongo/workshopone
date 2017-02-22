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
	for i := 0; i < 10; i++ {
		for key := range ages {
			fmt.Print(key, ":", ages[key], ", ")
		}
		fmt.Println()
	}
	// END OMIT
}
