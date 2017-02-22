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
	age, ok := ages["Baby"]
	if ok {
		fmt.Println("Baby age is", age)
	} else {
		fmt.Println("Name not in map")
	}
	// END OMIT
}
