package main

import (
	"fmt"
)

func main() {
	// START OMIT
	type Meters float64
	bob := Meters(1.73)
	whiskers := Meters(0.25)
	fmt.Println(bob / whiskers)
	// END OMIT
}
