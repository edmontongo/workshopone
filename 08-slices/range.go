package main

import (
	"fmt"
	"strings"
)

func main() {
	// BEGIN OMIT
	words := []string{"extra", "extra", "read", "all", "about", "it"}

	for i := range words { // HL
		fmt.Print(strings.ToUpper(words[i]), " ")
	}
	// END OMIT
}
