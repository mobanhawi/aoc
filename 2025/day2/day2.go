package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Solve https://adventofcode.com/2025/day/2
*/
func Solve() {
	input, err := os.ReadFile("2025/day2/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), ",")
	fmt.Println("2025/day/2", solve(lines))
}

func hasRepeatingPattern(s string) bool {
	n := len(s)
	// Try all possible pattern lengths from 1 to n/2
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen == 0 {
			pattern := s[:patternLen]
			matched := true
			// Check if all chunks match the first pattern
			for i := patternLen; i < n; i += patternLen {
				if s[i:i+patternLen] != pattern {
					matched = false
					break
				}
			}
			if matched {
				return true
			}
		}
	}
	return false
}

func solve(lines []string) int {
	sum := 0

	for _, line := range lines {
		numbers := strings.Split(line, "-")
		if len(numbers) != 2 {
			log.Fatalf("invalid instruction: %s", line)
		}
		start, _ := strconv.Atoi(numbers[0])
		end, _ := strconv.Atoi(numbers[1])
		for i := start; i <= end; i++ {
			stringNumber := strconv.Itoa(i)
			if hasRepeatingPattern(stringNumber) {
				sum += i
			}
		}
	}

	return sum
}
