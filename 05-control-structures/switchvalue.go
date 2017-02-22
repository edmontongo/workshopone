package main

import "fmt"

// START OMIT
func getFruitColor(fruit string) string {
	switch fruit {
	case "banana":
		return "yellow"
	case "orange":
		return "orange (duh)"
	case "apple", "strawberry":
		return "red"
	default:
		return "is too fancy"
	}
}

func main() {
	fruit := "orange"
	fmt.Println(fruit, ":", getFruitColor(fruit))
}

// END OMIT
