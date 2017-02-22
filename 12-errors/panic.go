package main

import "errors"

func main() {
	panic(errors.New("Unwilling to write hello world")) // HL
	println("Hello, World")
}
