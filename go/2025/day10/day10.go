package day9

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/mobanhawi/aoc/2025/util"
)

/*
Solve https://adventofcode.com/2025/day/10
*/
func Solve() {
	fmt.Println("2025/day/9 pt1", solvePt1(util.ReadLines("./2025/day10/input.txt")))
	fmt.Println("2025/day/9 pt2", solvePt2(util.ReadLines("./2025/day10/input.txt")))
}

type Vector []int

// Distance computes the Manhattan distance between two vectors
func (v Vector) Distance(o Vector) int {
	r := 0
	for i := range len(v) {
		r += v[i] - o[i]
	}
	return r
}

// dist manhattan distance between two points
func dist(p1, p2 []int) int {
	d := 0
	for i := range p1 {
		d += util.Abs(p1[i] - p2[i])
	}
	return d
}

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) int {
	for _, line := range lines {
		r := []rune(line)
		startTarget := slices.Index(r, '[')
		endTarget := slices.Index(r, ']')
		target := make(Vector, endTarget-startTarget-1)
		for i := startTarget + 1; i < endTarget; i++ {
			if r[i] == '#' {
				target[i-startTarget-1] = 1
			}
		}
		startJolts := slices.Index(r, '{')
		switches := parseSwitches(line[endTarget+2 : startJolts+2])
		fmt.Println("targets:", target)
		fmt.Println("switches:", switches)
		start := make([]int, len(target))
		count := 0
		minDist := -1
		for target.Distance(start) == 0 {
			for _, sw := range switches {
				newPos := make(Vector, len(start))
				for i := range start {
					if count%2 == 0 {

					}
					newPos[i] = start[i] + sw[i]
				}
				dist := target.Distance(newPos)
				if minDist == -1 || dist < minDist {
					minDist = dist
					start = newPos
				}
			}
			count++

		}

	}
	return 0
}

func solvePt2(lines []string) int {
	return 0
}

func parseSwitches(line string) []Vector {
	// Regex to match both single numbers (x) and pairs (x,y)
	re := regexp.MustCompile(`\((\d+)(?:,(\d+))?\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return []Vector{}
	}

	result := make([]Vector, 0, len(matches))

	for _, match := range matches {
		indices := make(Vector, 0, 2)
		maxIdx := 0

		// First number is always present
		idx1, _ := strconv.Atoi(match[1])
		indices = append(indices, idx1)
		if idx1 > maxIdx {
			maxIdx = idx1
		}

		// Check if there's a second number (pair)
		if match[2] != "" {
			idx2, _ := strconv.Atoi(match[2])
			indices = append(indices, idx2)
			if idx2 > maxIdx {
				maxIdx = idx2
			}
		}

		// Create array with all zeros
		arr := make([]int, maxIdx+1)

		// Set specified indices to 1
		for _, idx := range indices {
			arr[idx] = 1
		}

		result = append(result, arr)
	}

	return result
}
