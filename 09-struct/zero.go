package main

import (
	"fmt"
	"bytes"
	"strconv"
)

func main() {
	// START OMIT
	buf := &bytes.Buffer{}
	for i := 0; i < 10; i++ {
		buf.WriteString(strconv.Itoa(i))
	}
	fmt.Println(buf)
	// END OMIT
}
