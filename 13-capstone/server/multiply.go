package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var _ = ProblemGenerator(&MultiplyProblemGenerator{})

type MultiplyProblemGenerator struct{}

func (gen *MultiplyProblemGenerator) Generate() Problem {
	size := 3 + rand.Intn(4)
	mp := MultiplyProblem{}
	for i := 0; i < size; i++ {
		mp = append(mp, rand.Intn(1<<4)+1)
	}
	return mp
}

func (gen *MultiplyProblemGenerator) Documentation(wr io.Writer) {
	fmt.Fprint(wr,
		`Doing a GET with a token will provide 3-7 positive integers encoded in a JSON
array. To proceed to the next task, you must POST the numbers multiplied
together.`)
}

type MultiplyProblem []int

func (mp MultiplyProblem) Send(wr io.Writer) {
	enc := json.NewEncoder(wr)
	if err := enc.Encode(mp); err != nil {
		log.Println("Failed to encode multiply problem:", err.Error())
	}
}

func (mp MultiplyProblem) Validate(r io.Reader) bool {
	answer := 1
	for _, coefficient := range mp {
		answer *= coefficient
	}
	b := make([]byte, 32)
	n, err := r.Read(b)
	switch {
	case n > 16:
	case err != io.EOF && err != nil:
		log.Println("Failed to parse user's number: ", err.Error())
	case string(b[0:n]) != fmt.Sprint(answer):
	default:
		return true
	}
	return false
}
