package main

import (
	"fmt"
	"time"
)

// EMPTY_BEGIN OMIT
type empty interface{}

// EMPTY_END OMIT

// DECLARE_BEGIN OMIT
type Stringer interface {
	ToString() string
}

// DECLARE_END OMIT

// IMPLEMENT_BEGIN OMIT
type Simple string

func (s Simple) ToString() string {
	return string(s)
}

type Complex struct {
	Field1 int
	Field2 string
	Field3 float32
}

func (c Complex) ToString() string {
	return fmt.Sprintf("f1: %v, f2: %v, f3: %v now: %v",
		c.Field1, c.Field2, c.Field3, time.Now().Format("15:04:05"))
}

type Composition struct {
	Complex
}

// IMPLEMENT_END OMIT

// USE_BEGIN OMIT
func main() {
	cplx := Complex{
		1,
		"complex",
		2.2,
	}
	composition := Composition{cplx}
	strs := []Stringer{
		Simple("simple"),
		cplx,
		composition,
	}

	printStrings(strs)
}

func printStrings(strs []Stringer) {
	for _, s := range strs {
		fmt.Println(s.ToString())
	}
}

// USE_END OMIT
