package day12

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var (
	lines  = util.ReadLines("input.txt")
	lines0 = util.ReadLines("input0.txt")
)

func Test_solvePt1(t *testing.T) {
	t.Run("input.txt", func(t *testing.T) {
		want := 512
		got := solvePt1(lines)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
}
