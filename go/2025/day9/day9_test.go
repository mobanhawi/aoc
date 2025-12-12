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
	t.Run("input0.txt", func(t *testing.T) {
		want := 50
		got := solvePt1(lines0)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		want := 4748826374
		got := solvePt1(lines)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
}

func Test_solvePt2(t *testing.T) {
	t.Run("input0.txt", func(t *testing.T) {
		want := 24
		got := solvePt2(lines0)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		want := 1554370486
		got := solvePt2(lines)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
}
