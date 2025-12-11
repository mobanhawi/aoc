package day10

import (
	"testing"

	"github.com/mobanhawi/aoc/2025/util"
)

var l = util.ReadLines("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day9
// cpu: Apple M1 Pro
// BenchmarkSolutionPt2
// BenchmarkSolutionPt2-8   	      55	  21025236 ns/op
// BenchmarkSolutionPt1
// BenchmarkSolutionPt1-8   	     190	   6222427 ns/op
func BenchmarkSolutionPt2(t *testing.B) {
	for t.Loop() {
		solvePt2(l)
	}
}
func BenchmarkSolutionPt1(t *testing.B) {
	for t.Loop() {
		solvePt1(l)
	}
}
