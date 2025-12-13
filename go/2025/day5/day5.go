package day5

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mobanhawi/aoc/util"
)

/*
Solve https://adventofcode.com/2025/day/5
*/
func Solve() {
	fmt.Println("2025/day/5 pt1", solvePt1(util.ReadLines("./2025/day5/input.txt")))
	fmt.Println("2025/day/5 pt2", solvePt2(util.ReadLines("./2025/day5/input.txt")))
}

type Range struct {
	min int
	max int
}

// solvePt2 solves part 1 of the puzzle
// It counts how many numbers are valid according to the given ranges
func solvePt2(lines []string) int {
	ranges := make([]Range, 0)
	validIDs := 0
	for _, line := range lines {
		if strings.Contains(line, "-") {
			var r Range
			fmt.Sscanf(line, "%d-%d", &r.min, &r.max)
			ranges = append(ranges, r)
		}
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.min - b.min
	})

	// combine overlapping ranges
	combined := make([]Range, 0)
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].min <= current.max {
			if ranges[i].max > current.max {
				current.max = ranges[i].max
			}
		} else {
			combined = append(combined, current)
			current = ranges[i]
		}
	}
	combined = append(combined, current)

	// count valid IDs
	for _, r := range combined {
		validIDs += r.max - r.min + 1
	}

	return validIDs
}

// solvePt1 solves part 1 of the puzzle
// It counts how many numbers are valid according to the given ranges
func solvePt1(lines []string) int {
	ranges := make([]Range, 0)
	numbers := make([]int, 0)
	validIDs := 0
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.Contains(line, "-") {
			var r Range
			fmt.Sscanf(line, "%d-%d", &r.min, &r.max)
			ranges = append(ranges, r)
		} else {
			var num int
			fmt.Sscanf(line, "%d", &num)
			numbers = append(numbers, num)
		}
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.min - b.min
	})

	for _, num := range numbers {
		if isValid(num, ranges) {
			validIDs++
		}
	}

	return validIDs
}

// isValid checks if a number is valid according to the given ranges
func isValid(num int, ranges []Range) bool {
	for _, r := range ranges {
		if num >= r.min && num <= r.max {
			return true
		}
	}
	return false
}
