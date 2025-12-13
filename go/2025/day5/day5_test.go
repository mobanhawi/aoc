package day5

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

func Test_solvePt1(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		lines := util.ReadLines("input.txt")
		want := 865
		got := solvePt1(lines)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
}

func Test_solvePt2(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		lines := util.ReadLines("input.txt")
		want := 352556672963116
		got := solvePt2(lines)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
}
