package main

import "fmt"

func main() {
	// BEGIN OMIT
	var s string = "Hi everybody"
	fmt.Println(s)
	var raw string = `Hi "everybody"`
	fmt.Println(raw)
	// END OMIT
}
