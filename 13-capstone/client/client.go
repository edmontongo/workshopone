package main

import (
	"flag"
	"log"
)

var (
	name = flag.String("name", "", "Your name to be supplied to the server")
	url  = flag.String("url", "http://localhost:8080", "URL of workshop server")
)

func main() {
	flag.Parse()
	if *name == "" {
		log.Fatal("Name flag is required!")
	}

}
