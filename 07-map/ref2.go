package main

import "fmt"

func main() {
	ref := map[int]int{}
	modify(ref)
	// START OMIT
	var other = ref
	other[1] *= 42
	fmt.Println(other)
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
