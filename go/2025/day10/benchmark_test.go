package day10

import (
	"testing"

	"github.com/mobanhawi/aoc/util"
)

var l = util.ReadLines("input.txt")

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day10
// cpu: Apple M1 Pro
// BenchmarkSolutionPt1
// BenchmarkSolutionPt1-8   	     139	   8541663 ns/op
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
