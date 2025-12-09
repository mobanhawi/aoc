package day9

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

var (
	lines  = util.ReadLines("input.txt")
	lines0 = util.ReadLines("input0.txt")
)

func Test_solvePt1(t *testing.T) {
	want0 := 50
	want1 := 4748826374
	got1 := solvePt1(lines)
	got0 := solvePt1(lines0)

	if got0 != want0 {
		t.Errorf("Test solvePt1 wanted %v got %v", want0, got0)
	}
	if got1 != want1 {
		t.Errorf("Test solvePt1 wanted %v got %v", want1, got1)
	}
}

func Test_solvePt2(t *testing.T) {
	want0 := 0
	want1 := 0
	got1 := solvePt2(lines)
	got0 := solvePt2(lines0)

	if got0 != want0 {
		t.Errorf("Test solvePt2 wanted %v got %v", want0, got0)
	}
	if got1 != want1 {
		t.Errorf("Test solvePt2 wanted %v got %v", want1, got1)
	}
}
