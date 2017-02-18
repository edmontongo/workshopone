package main

import (
	"fmt"
	"io"
	"math/rand"
)

var _ = ProblemGenerator(&SquareProblemGenerator{})

type SquareProblemGenerator struct{}

func (gen *SquareProblemGenerator) Generate() Problem {
	return SquareProblem{rand.Intn(1 << 15)}
}

func (gen *SquareProblemGenerator) Documentation(wr io.Writer) {
}

type SquareProblem struct {
	square int
}

func (sp SquareProblem) Send(wr io.Writer) {
	fmt.Fprint(wr, sp.square)
}

func (sp SquareProblem) Validate(r io.Reader) bool {
	b := make([]byte, 32)
	n, err := r.Read(b)
	switch {
	case err != nil:
	case n > 16:
	case string(b[0:n]) != fmt.Sprint(sp.square*sp.square):
	default:
		return true
	}
	return false
}
