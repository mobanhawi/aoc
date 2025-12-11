package day10

import (
	"container/list"
	"fmt"
	"math/bits"
	"regexp"
	"slices"
	"strconv"

	"github.com/mobanhawi/aoc/2025/util"
)

/*
Solve https://adventofcode.com/2025/day/10
*/
func Solve() {
	fmt.Println("2025/day/10 pt1", solvePt1(util.ReadLines("./2025/day10/input.txt")))
	fmt.Println("2025/day/10 pt2", solvePt2(util.ReadLines("./2025/day10/input.txt")))
}

type State struct {
	Position Vector
	Count    int
}

type Vector uint

// Distance computes the Manhattan distance between two vectors
func (v Vector) Distance(o Vector) int {
	// Count the number of differing bits by XORing and counting set bits
	return bits.OnesCount(uint(v ^ o))
}

func (v Vector) String() string {
	return fmt.Sprintf("%b", v)
}

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) int {
	total := 0
	for _, line := range lines {
		r := []rune(line)
		startTarget := slices.Index(r, '[')
		endTarget := slices.Index(r, ']')
		bitWidth := endTarget - startTarget - 1
		var target Vector = 0
		for i := startTarget + 1; i < endTarget; i++ {
			if r[i] == '#' {
				target = target | (1 << (bitWidth - 1 - (i - startTarget - 1)))
			}
		}
		startJolts := slices.Index(r, '{')
		switches := parseSwitches(line[endTarget+2:startJolts+2], bitWidth)
		fmt.Println("targets:", target)
		fmt.Println("switches:", switches)
		var currentPos Vector = 0
		l := list.New()
		l.PushBack(State{Position: currentPos, Count: 0})
		visited := util.NewSet[Vector]()
		count := 0

		for l.Len() > 0 {
			state := l.Remove(l.Front()).(State)
			currentPos = state.Position
			count = state.Count
			if currentPos == target {
				fmt.Println("Reached target:", target, "in", count, "switches")
				total += count
				break
			}
			if visited.Contains(currentPos) {
				continue
			}
			visited.Add(currentPos)
			for _, sw := range switches {
				newPos := currentPos ^ sw // toggle switch using XOR
				l.PushBack(State{Position: newPos, Count: count + 1})
			}
		}
		if currentPos != target {
			panic("Could not reach target: " + target.String())
		}
	}
	return total
}

func solvePt2(lines []string) int {
	return 0
}

func parseSwitches(line string, bitWidth int) []Vector {
	// Regex to match parentheses containing comma-separated numbers: (x), (x,y), (x,y,z), etc.
	re := regexp.MustCompile(`\(([0-9,]+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return []Vector{}
	}

	result := make([]Vector, 0, len(matches))

	for _, match := range matches {
		var v Vector = 0
		// Split the comma-separated indices
		indices := regexp.MustCompile(`\d+`).FindAllString(match[1], -1)

		for _, idxStr := range indices {
			idx, _ := strconv.Atoi(idxStr)
			v = v | (1 << (bitWidth - 1 - idx)) // MSB/LSB flipped
		}
		result = append(result, v)
	}

	return result
}
