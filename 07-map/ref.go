package main

import "fmt"

func main() {
	// START OMIT
	ref := map[int]int{}
	modify(ref)
	fmt.Println(ref)
	// END OMIT
}

// FUNC_S OMIT
func modify(ref map[int]int) {
	ref[1] = 1
	ref[2] = 4
	ref[3] = 9
	ref[4] = 16
}

// FUNC_E OMIT
