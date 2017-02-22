package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMultiplyProblemWrite(t *testing.T) {
	buf := bytes.Buffer{}
	mp := MultiplyProblem{1, 2, 3, 4}
	mp.Send(&buf)
	line, _ := buf.ReadString(0)
	if strings.TrimSpace(line) != "[1,2,3,4]" {
		t.Error("Expected '[1,2,3,4]' to be written, got:", line)
	}
}

func TestMultiplyProblemValidatePasses(t *testing.T) {
	mp := MultiplyProblem{1, 2, 3, 4}

	if !mp.Validate(bytes.NewBufferString("24")) {
		t.Error("Expect '24' to be correct")
	}
}

func TestMultiplyProblemValidateFails(t *testing.T) {
	mp := MultiplyProblem{1, 2, 3, 4}
	inputs := []string{"an orange", "64", "24.0"}

	for _, test := range inputs {
		if mp.Validate(bytes.NewBufferString(test)) {
			t.Errorf("Did not expect '%s' to be valid input", test)
		}
	}
}
