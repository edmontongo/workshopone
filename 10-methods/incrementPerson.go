package main

import "fmt"

type Incrementor struct {
	int
}

func (i Incrementor) Print() {
	fmt.Printf("The incrementor value is %d.\n", i.int)
}

func (i *Incrementor) Increment(amount int) {
	i.int += amount
}

// BEGIN OMIT

type Person struct {
	Age  Incrementor
	Name string
}

func (p Person) print() {
	fmt.Printf("%s is %d years old.\n", p.Name, p.Age.int)
}

func main() {
	person := Person{
		Age:  Incrementor{15},
		Name: "Strans",
	}
	person.Age.Increment(3)
	person.print()
}

// END OMIT
