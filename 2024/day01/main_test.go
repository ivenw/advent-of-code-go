package main

import (
	"testing"
)

const input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPartOne(t *testing.T) {
	data := parse([]byte(input))
	answer := partOne(data)
	want := 11
	if answer != want {
		t.Errorf("Expected: %v, Got: %v", want, answer)
	}
}

func TestPartTwo(t *testing.T) {
	data := parse([]byte(input))
	answer := partTwo(data)
	want := 31
	if answer != want {
		t.Errorf("Expected: %v, Got: %v", want, answer)
	}
}