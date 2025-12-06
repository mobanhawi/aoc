package day2

import (
	"os"
	"strings"
	"testing"
)

// goos: darwin
// goarch: arm64
// pkg: github.com/mobanhawi/aoc/2025/day2
// cpu: Apple M1 Pro
// BenchmarkSolution
// BenchmarkSolution-8           16          66947888 ns/op
func BenchmarkSolution(t *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	l := strings.Split(strings.TrimSpace(string(input)), ",")
	for t.Loop() {
		solve(l)
	}
}
