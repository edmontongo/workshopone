package main

import "fmt"

func main() {
	// START OMIT
	ref := map[int]int{}
	var other = ref
	other[3] = 42
	fmt.Println(ref)
	// END OMIT
}

