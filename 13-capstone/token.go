package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// BEGIN OMIT
	if resp, err := http.Get("http://workshop.bellstone.ca/task/new?name=SlideDeck"); err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, resp.Body)
	}
	// END OMIT
}
