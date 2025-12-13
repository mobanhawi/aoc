package day8

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var l = util.ReadLines("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day8
// cpu: Apple M1 Pro
// BenchmarkSolutionPt2
// BenchmarkSolutionPt2-8   	      25	  45465247 ns/op
// BenchmarkSolutionPt1
// BenchmarkSolutionPt1-8   	      27	  41038557 ns/op
func BenchmarkSolutionPt2(t *testing.B) {
	for t.Loop() {
		solvePt2(l)
	}
}
func BenchmarkSolutionPt1(t *testing.B) {
	for t.Loop() {
		solvePt1(l, 1000)
	}
}
