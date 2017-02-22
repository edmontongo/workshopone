package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

var _ = ProblemGenerator(&FrequencyProblemGenerator{})

type ParagraphFetcher interface {
	FetchParagraph() string
}

type FrequencyProblemGenerator struct {
	ParagraphFetcher
}

func (gen *FrequencyProblemGenerator) Generate() Problem {
	p := gen.FetchParagraph()
	counts := map[string]int{}
	for _, word := range strings.Split(p, " ") {
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

	return FrequencyProblem{p, [3]string{topCounts[0].word, topCounts[1].word, topCounts[2].word}}
}

func (gen *FrequencyProblemGenerator) Documentation(wr io.Writer) {
	fmt.Fprint(wr,
		`Doing a GET with a token will provide a paragraph, stripped of punctuation
and capitalization. To proceed to the next task, you must POST with the
three most common words in that paragraph separated by commas. Order does
not matter.

strings.Split will likely be helpful.
`)
}

type FrequencyProblem struct {
	paragraph string
	words     [3]string
}

func (fp FrequencyProblem) Send(wr io.Writer) {
	fmt.Fprint(wr, fp.paragraph)
}

func (fp FrequencyProblem) Validate(r io.Reader) bool {
	bufSize := 128
	b := make([]byte, bufSize)
	n, err := r.Read(b)
	switch err {
	case io.EOF:
	case nil:
	default:
		log.Println("Failed to parse user's input: ", err.Error())
		return false
	}

	result := strings.Split(string(b[:n]), ",")
	if len(result) != 3 {
		return false
	}
	foundWords := map[string]struct{}{}
	for _, word := range result {
		foundWords[word] = struct{}{}
	}
	for _, word := range fp.words {
		if _, ok := foundWords[word]; !ok {
			return false
		}
	}
	return true
}
