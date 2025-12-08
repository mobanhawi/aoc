package day2

import (
	"os"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	want := 24774350322
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), ",")
	got := solve(lines)
	if got != want {
		t.Errorf("Test solve wanted %v got %v", want, got)
	}
}
