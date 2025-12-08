package day4

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

var l = util.ReadLines("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day4
// cpu: Apple M1 Pro
// BenchmarkSolution
// BenchmarkSolution-8   	      19	  58722588 ns/op
func BenchmarkSolution(t *testing.B) {
	for t.Loop() {
		solve(l)
	}
}
