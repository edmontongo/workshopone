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
	if _, ok := ages["Baby"]; ok {
		fmt.Println("Baby age is known")
	}
	// END OMIT
}
