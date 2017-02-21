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
	ages["Baby"] += 3
	ages["Baby2"] += 2
	fmt.Println(ages)
	// END OMIT
}
