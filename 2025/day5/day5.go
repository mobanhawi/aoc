package day5

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
Solve https://adventofcode.com/2025/day/5
*/
func Solve() {
	input, err := os.ReadFile("2025/day5/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println("2025/day/5", solvePt1(lines))
}

type Range struct {
	min int
	max int
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
