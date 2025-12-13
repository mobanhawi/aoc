package day11

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var (
	lines   = util.ReadLines("input.txt")
	lines0  = util.ReadLines("input0.txt")
	lines00 = util.ReadLines("input00.txt")
)

func Test_solvePt1(t *testing.T) {
	t.Run("input0.txt", func(t *testing.T) {
		want := 5
		got := solvePt1(lines0)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		want := 796
		got := solvePt1(lines)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
}

func Test_solvePt2(t *testing.T) {
	t.Run("input00.txt", func(t *testing.T) {
		want := 2
		got := solvePt2(lines00)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		want := 294053029111296
		got := solvePt2(lines)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
}
