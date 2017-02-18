package main

import "io"

type ProblemGenerator interface {
	Generate() Problem
	Documentation(wr io.Writer)
}

type Problem interface {
	Send(wr io.Writer)
	Validate(r io.Reader) bool
}
