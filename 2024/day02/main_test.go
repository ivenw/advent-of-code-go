package main

import (
	"testing"
)

const input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPartOne(t *testing.T) {
	data := parse(input)
	answer := partOne(data)
	want := 2
	if answer != want {
		t.Errorf("Expected: %v, Got: %v", want, answer)
	}
}

func TestPartTwo(t *testing.T) {
	data := parse(input)
	answer := partTwo(data)
	want := 4
	if answer != want {
		t.Errorf("Expected: %v, Got: %v", want, answer)
	}
}
