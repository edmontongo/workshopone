package main

import (
	"fmt"
)

// BEGIN OMIT
func main() {
	ñümβęr := 4
	fmt.Println("ñümβęr is", ñümβęr)

	šqúàrë(&ñümβęr)
	fmt.Println("After šqúàrë, ñümβęr is", ñümβęr)
}

func šqúàrë(ptr *int) {
	*ptr = *ptr * *ptr
}

// END OMIT
