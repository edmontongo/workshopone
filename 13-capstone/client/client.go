package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	squareProblem(token)
	frequencyProblem(token)
}

func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func exitOnBadResponse(resp *http.Response, err error) {
	exitOnError(err)
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
}

func mustGetString(url string) string {
	resp, err := http.DefaultClient.Get(url)
	exitOnBadResponse(resp, err)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	exitOnError(err)
	return string(b)
}

func mustPostString(url, data string) {
	exitOnBadResponse(http.DefaultClient.Post(url, "", bytes.NewBufferString(data)))
}

func getToken() string {
	token := mustGetString(*url + "/task/new?name=" + *name)
	mustPostString(*url+"/task/"+token, "")
	return token
}

func squareProblem(token string) {
	url := *url + "/task/square/" + token
	number, err := strconv.Atoi(mustGetString(url))
	exitOnError(err)
	number = number * number
	mustPostString(url, strconv.Itoa(number))
}

func frequencyProblem(token string) {
	url := *url + "/task/frequency/" + token
	paragraph := mustGetString(url)

	counts := map[string]int{}
	for _, word := range strings.Split(paragraph, " ") {
		counts[word] = counts[word] + 1
	}
	topCounts := [3]struct {
		word  string
		count int
	}{}
	for word, count := range counts {
		for i, tc := range topCounts {
			if tc.count > count {
				continue
			}
			if i > 0 {
				topCounts[i-1] = tc
			}
			topCounts[i].word = word
			topCounts[i].count = count
		}
	}

	mustPostString(url, strings.Join([]string{topCounts[0].word, topCounts[1].word, topCounts[2].word}, ","))
}
