package main

import "fmt"

type Incrementor struct {
	int
}

// BEGIN OMIT

func (i *Incrementor) Print() *Incrementor {
	fmt.Printf("The incrementor value is %d.\n", i.int)
	return i
}

func (i *Incrementor) Increment(amount int) *Incrementor {
	i.int += amount
	return i
}

func main() {
	inc := &Incrementor{}
	inc.Print().Increment(1).Increment(5)
	inc.Print()
}

// END OMIT
