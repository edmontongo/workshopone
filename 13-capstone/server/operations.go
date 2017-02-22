package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var _ = ProblemGenerator(&OperationsProblemGenerator{})

type OperationsProblemGenerator struct{}

func (gen *OperationsProblemGenerator) Generate() Problem {
	size := 50 + rand.Intn(50)
	mp := OperationsProblem{}
	for i := 0; i < size; i++ {
		mp = append(mp, NewOperation())
	}
	return mp
}

func (gen *OperationsProblemGenerator) Documentation(wr io.Writer) {
	fmt.Fprint(wr,
		`Doing a GET with a token will provide 3-7 positive integers encoded in a JSON
array. To proceed to the next task, you must POST the numbers multiplied
together.`)
}

type op string

const (
	add = op("add")
	sub = op("sub")
	mul = op("mul")
	pow = op("pow")
)

type Operation struct {
	Op    op  `json:"op"`
	Term1 int `json:"first"`
	Term2 int `json:"second"`
}

func (o Operation) Solve() int {
	switch o.Op {
	case add:
		return o.Term1 + o.Term2
	case sub:
		return o.Term1 - o.Term2
	case mul:
		return o.Term1 * o.Term2
	case pow:
		return int(math.Pow(float64(o.Term1), float64(o.Term2)))
	default:
		log.Panic("Invalid op type: ", string(o.Op))
	}
	return -1
}

var ops = map[int]op{
	0: add,
	1: sub,
	2: mul,
	3: pow,
}

func NewOperation() Operation {
	randomOp, ok := ops[rand.Intn(4)]
	if !ok {
		log.Panic("Unknown operation seclected!")
	}
	o := Operation{Op: randomOp}
	switch o.Op {
	case add:
		fallthrough
	case sub:
		o.Term1 = rand.Intn(1<<30) - 1<<29
		o.Term2 = rand.Intn(1<<30) - 1<<29
	case mul:
		o.Term1 = rand.Intn(1<<15) - 1<<14
		o.Term2 = rand.Intn(1<<15) - 1<<14
	case pow:
		o.Term1 = rand.Intn(21) - 10
		o.Term2 = rand.Intn(9) + 1
	default:
	}
	return o
}

type OperationsProblem []Operation

func (opp OperationsProblem) Send(wr io.Writer) {
	enc := json.NewEncoder(wr)
	if err := enc.Encode(opp); err != nil {
		log.Panic("Failed to encode Operations problem:", err.Error())
	}
}

func (opp OperationsProblem) Validate(r io.Reader) bool {
	answers := []int{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&answers); err != nil {
		return false
	}
	if len(answers) != len(opp) {
		return false
	}
	for i, o := range opp {
		if o.Solve() != answers[i] {
			return false
		}
	}
	return true
}
