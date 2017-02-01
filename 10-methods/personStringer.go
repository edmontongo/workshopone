package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

// BEGIN OMIT

func (p Person) print() {
	fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
}

// END OMIT

func main() {
	p := Person{Age: 15, Name: "Strans"}
	fmt.Println(p)
}
