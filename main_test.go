package main

import "testing"

func Test_SolvePart1(t *testing.T) {
	result, err := SolvePart1("./example")
	if err != nil {
		t.Error(" %w", err)
	}
	if result != 5 {
		t.Error("result wasn't correct, wanted 5 but received", result)
	}
}

func Test_SolvePart1_RealAnswer(t *testing.T) {
	result, err := SolvePart1("./input")
	if err != nil {
		t.Error(" %w", err)
	}
	if result != 7864 {
		t.Error("result wasn't correct, wanted 7864 but received", result)
	}
}
