package day4

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

func Test_solve(t *testing.T) {
	lines := util.ReadLines("input.txt")
	want := 8727
	got := solve(lines)

	if got != want {
		t.Errorf("Test solve wanted %v got %v", want, got)
	}

}
