package day12

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var l = util.ReadLines("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day12
// cpu: Apple M1 Pro
// BenchmarkSolutionPt1
// BenchmarkSolutionPt1-8   	    1887	    641951 ns/op
func BenchmarkSolutionPt1(t *testing.B) {
	for t.Loop() {
		solvePt1(l)
	}
}
