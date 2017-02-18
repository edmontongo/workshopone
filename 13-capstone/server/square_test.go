package main

import (
	"bytes"
	"testing"
)

func TestSquareProblemWrite(t *testing.T) {
	buf := bytes.Buffer{}
	sp := SquareProblem{75}
	sp.Send(&buf)
	line, _ := buf.ReadString(0)
	if line != "75" {
		t.Error("Expected '75' to be written, got:", line)
	}
}

func TestSquareProblemValidatePasses(t *testing.T) {
	sp := SquareProblem{4}

	if !sp.Validate(bytes.NewBufferString("16")) {
		t.Error("Expect '16' to be correct")
	}
}

func TestSquareProblemValidateFails(t *testing.T) {
	sp := SquareProblem{4}
	inputs := []string{"an orange", "64", "16.0"}

	for _, test := range inputs {
		if sp.Validate(bytes.NewBufferString(test)) {
			t.Errorf("Did not expect '%s' to be valid input", test)
		}
	}
}
