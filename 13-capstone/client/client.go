package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	token := getToken()
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getToken() string {
	resp, err := http.DefaultClient.Get(*url + "/task/new?name=" + *name)
	exitOnError(err)
	defer resp.Body.Close()
	var token string
	_, err = fmt.Fscanf(resp.Body, "%s", &token)
	exitOnError(err)
	return token
}
