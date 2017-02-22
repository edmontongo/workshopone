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
	Incrementor
	Name string
}

func (p Person) print() {
	fmt.Printf("%s is %d years old.\n", p.Name, p.int)
}

func main() {
	person := Person{
		Incrementor: Incrementor{15},
		Name:        "Strans",
	}
	person.Increment(3)
	person.print()
}

// END OMIT
