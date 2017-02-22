package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("oops!")

	fmt.Printf("Error: %+v = %s = %s\n", e, e, e.Error())

	e = fmt.Errorf("I am also an error: %s", e)
	fmt.Println("Annotated:", e)
}
