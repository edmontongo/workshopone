package main

import "fmt"

// BEGIN OMIT

func printPerson(age int, name string) {
	fmt.Printf("%s is %d years old.\n", name, age)
}

func main() {
	printPerson(15, "Strans")
}

// END OMIT
