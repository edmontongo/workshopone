package main

import (
	"fmt"
	"net/http"
)

func main() {
	if resp, err := http.Get("http://workshop.bellstone.ca/task/"); err != nil {
		fmt.Println("Error!", err.Error())
	} else {
		fmt.Println(resp.Status)
	}
}
