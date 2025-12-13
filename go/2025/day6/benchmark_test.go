package day6

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var l = util.ReadLines("input.txt")
var l2 = util.ReadLinesNoTrim("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day6
// cpu: Apple M1 Pro
// BenchmarkSolutionPt2
// BenchmarkSolutionPt2-8   	   22183	     53780 ns/op
// BenchmarkSolutionPt1
// BenchmarkSolutionPt1-8   	    6370	    183658 ns/op
func BenchmarkSolutionPt2(t *testing.B) {
	for t.Loop() {
		solvePt2(l2)
	}
}
func BenchmarkSolutionPt1(t *testing.B) {
	for t.Loop() {
		solvePt1(l)
	}
}
