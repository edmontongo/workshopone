package main

import "fmt"

func printPerson(age int, name string) {
	fmt.Printf("%s is %d years old.\n", name, age)
}

// BEGIN OMIT

type Person struct {
	Age  int
	Name string
}

func main() {
	p := Person{Age: 15, Name: "Strans"}
	printPerson(p.Age, p.Name)
}

// END OMIT
