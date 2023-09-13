package main

import "testing"

func Test_SolvePart1(t *testing.T) {
	result, err := SolvePart1("./example")
	if err != nil {
		t.Error(" %w", err)
	}
	wanted := 5
	if result != wanted {
		t.Error("result wasn't correct, wanted", wanted, "but received", result)
	}
}

func Test_SolvePart1_RealAnswer(t *testing.T) {
	result, err := SolvePart1("./input")
	if err != nil {
		t.Error(" %w", err)
	}
	wanted := 7864
	if result != wanted {
		t.Error("result wasn't correct, wanted", wanted, "but received", result)
	}
}

func Test_SolvePart2(t *testing.T) {
	result, err := SolvePart2("./example")
	if err != nil {
		t.Error(" %w", err)
	}
	wanted := 4
	if result != wanted {
		t.Error("result wasn't correct, wanted", wanted, "but received", result)
	}
}

func Test_SolvePart2_RealAnswer(t *testing.T) {
	result, err := SolvePart2("./input")
	if err != nil {
		t.Error(" %w", err)
	}
	wanted := 1695
	if result != wanted {
		t.Error("result wasn't correct, wanted", wanted, "but received", result)
	}
}
