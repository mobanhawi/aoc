package day6

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

func Test_solve(t *testing.T) {
	lines := util.ReadLines("input.txt")
	want1 := 4449991244405
	want2 := 3263827
	got1 := solvePt1(lines)
	got2 := solvePt2(util.ReadLines("input0.txt"))

	if got1 != want1 {
		t.Errorf("Test solvePt1 wanted %v got %v", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("Test solvePt2 wanted %v got %v", want2, got2)
	}
}
