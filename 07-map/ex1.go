package main

//for tests run go get golang.org/x/tour/wc
import (
	"golang.org/x/tour/wc"
	"strings"
)

// START OMIT
func WordCount(s string) map[string]int {
	count := map[string]int{}
	for _, word := range strings.Fields(s) {
		count[word]++
	}
	return count
}

// END OMIT

func main() {
	wc.Test(WordCount)
}
