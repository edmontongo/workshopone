package main

import "fmt"

// START OMIT
func getInterestRate(balance float32) float32 {
	var rate float32

	if balance <= 1000.0 { // HL
		rate = 0.005 // HL
	} // HL

	return rate
}

func main() {
	fmt.Println("rate:", getInterestRate(100.0))
}

// END OMIT
