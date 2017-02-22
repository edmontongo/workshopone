package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math"
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
	multiplyProblem(token)
	operationsProblem(token)
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

func mustGetJSON(url string, v interface{}) {
	resp, err := http.DefaultClient.Get(url)
	exitOnBadResponse(resp, err)
	defer resp.Body.Close()
	enc := json.NewDecoder(resp.Body)
	exitOnError(enc.Decode(v))
}

func mustPostString(url, data string) {
	exitOnBadResponse(http.DefaultClient.Post(url, "", bytes.NewBufferString(data)))
}

func mustPostJSON(url string, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	exitOnError(enc.Encode(v))
	exitOnBadResponse(http.DefaultClient.Post(url, "", buf))
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

func multiplyProblem(token string) {
	url := *url + "/task/multiply/" + token
	nums := []int{}
	mustGetJSON(url, &nums)
	answer := 1
	for _, n := range nums {
		answer *= n
	}
	mustPostString(url, strconv.Itoa(answer))
}

func operationsProblem(token string) {
	url := *url + "/task/operations/" + token
	ops := []struct {
		Op            string
		First, Second int
	}{}
	mustGetJSON(url, &ops)
	answers := []int{}
	for _, op := range ops {
		ans := 0
		switch op.Op {
		case "add":
			ans = op.First + op.Second
		case "sub":
			ans = op.First - op.Second
		case "mul":
			ans = op.First * op.Second
		case "pow":
			ans = int(math.Pow(float64(op.First), float64(op.Second)))
		default:
			panic(op.Op)
		}
		answers = append(answers, ans)
	}
	mustPostJSON(url, answers)
}
