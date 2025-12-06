package day1

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

func Test_solve(t *testing.T) {
	want := 5887
	got := solve(util.ReadLines("input.txt"), 50)
	if got != want {
		t.Errorf("Test solve wanted %v got %v", want, got)
	}
}
