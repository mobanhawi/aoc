package day3

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

func Test_solvePt1(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		lines := util.ReadLines("input.txt")
		want := 17301
		got := solvePt1(lines)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
}

func Test_solvePt2(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		lines := util.ReadLines("input.txt")
		want := 172162399742349
		got := solvePt2(lines)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
}
