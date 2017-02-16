package main

import (
	"fmt"
)

func main() {
	// START OMIT
	type Meters float64
	alice := int(2)
	bob := Meters(1.73)
	fmt.Println(alice + bob) // mismatched types

	b := []byte{'h', 'i'}
	s := string(b)
	b[1] = 'a'
	fmt.Printf("b=%q s=%q\n", b, s)
	// END OMIT
	_ = alice
	_ = bob
}
