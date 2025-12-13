package day1

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var l = util.ReadLines("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day1
// cpu: Apple M1 Pro
// BenchmarkSolution
// BenchmarkSolution-8         1870            638058 ns/op
func BenchmarkSolution(t *testing.B) {
	for t.Loop() {
		solve(l, 50)
	}
}
