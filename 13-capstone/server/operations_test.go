package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestOperationSolve(t *testing.T) {
	tests := []struct {
		o   Operation
		res int
	}{
		{Operation{add, 1, 2}, 3},
		{Operation{sub, 4, 5}, -1},
		{Operation{mul, 3, 6}, 18},
		{Operation{pow, 4, 7}, 16384},
	}

	for _, test := range tests {
		res := test.o.Solve()
		if res != test.res {
			t.Errorf("Failed to solve opersion %s, expected %d got %d", test.o.Op, test.res, res)
		}
	}
}

func TestOperationsProblemWrite(t *testing.T) {
	buf := bytes.Buffer{}
	opp := OperationsProblem{{add, 1, 2}, {pow, 9, 8}}
	opp.Send(&buf)
	line, _ := buf.ReadString(0)
	line = strings.TrimSpace(line)
	if line != `[{"op":"add","first":1,"second":2},{"op":"pow","first":9,"second":8}]` {
		t.Error(`Expected '[{"op":"add","first":1,"second":2},{"op","pow","first":9,"second":8}]' to be written, got:`, line)
	}
}

func TestOperationsProblemValidatePasses(t *testing.T) {
	opp := OperationsProblem{{add, 1, 2}, {pow, 9, 8}}

	if !opp.Validate(bytes.NewBufferString("[3,43046721]")) {
		t.Error("Expect '[3,43046721]' to be correct")
	}
}

func TestOperationsProblemValidateFails(t *testing.T) {
	opp := OperationsProblem{{add, 1, 2}, {pow, 9, 8}}
	inputs := []string{"an orange", "64", "3,43046721", "[43046721,3]", "[3]", "[3,43046721,0]", "{3,43046721}"}

	for _, test := range inputs {
		if opp.Validate(bytes.NewBufferString(test)) {
			t.Errorf("Did not expect '%s' to be valid input", test)
		}
	}
}
