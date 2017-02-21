package main

import (
	"fmt"
	"strings"
)

// START OMIT
type LetterCase uint

const (
	LowerCase LetterCase = iota
	UpperCase
)

func Print(s string, letterCase LetterCase) {
	if letterCase == LowerCase {
		fmt.Println(strings.ToLower(s))
	} else {
		fmt.Println(strings.ToUpper(s))
	}
}

func main() {
	Print("hI", UpperCase)
}

// END OMIT
