package main

import (
	"fmt"
)

func main() {
// START OMIT
	feet := 5.68
	meters := 1.73
	fmt.Println(feet + meters) // Is the result meaningful?

// START DISTINCT OMIT
	type Feet float64
	type Meters float64
	mile := Feet(5280)
	km := Meters(1000)
	fmt.Println(mile + km) // type mismatch
// END DISTINCT OMIT
// END OMIT
}
