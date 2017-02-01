package main

import "fmt"

// BEGIN OMIT

type Incrementor struct {
	int
}

func (i Incrementor) Print() {
	fmt.Printf("The incrementor value is %d.\n", i.int)
}

func (i Incrementor) Increment(amount int) {
	i.int += amount
}

func main() {
	inc := Incrementor{}
	inc.Print()
	inc.Increment(1)
	inc.Increment(5)
	inc.Print()
}

// END OMIT
