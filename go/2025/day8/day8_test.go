package day8

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
		want := 40
		got := solvePt1(lines0, 10)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		want := 131580
		got := solvePt1(lines, 1000)
		if got != want {
			t.Errorf("Test solvePt1 wanted %v got %v", want, got)
		}
	})
}

func Test_solvePt2(t *testing.T) {
	t.Run("input0.txt", func(t *testing.T) {
		want := 25272
		got := solvePt2(lines0)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		want := 6844224
		got := solvePt2(lines)
		if got != want {
			t.Errorf("Test solvePt2 wanted %v got %v", want, got)
		}
	})
}
