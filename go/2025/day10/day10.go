package day9

import (
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
	fmt.Println("2025/day/9 pt1", solvePt1(util.ReadLines("./2025/day10/input.txt")))
	fmt.Println("2025/day/9 pt2", solvePt2(util.ReadLines("./2025/day10/input.txt")))
}

type Vector uint

// Distance computes the Manhattan distance between two vectors
func (v Vector) Distance(o Vector) int {
	// Count the number of differing bits by XORing and counting set bits
	//fmt.Println("v:", v, "o:", o, "v^o:", v^o)
	return bits.OnesCount(uint(v ^ o))
}

func (v Vector) String() string {
	return fmt.Sprintf("%b", v)
}

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) (switchCount int) {
	// count number of presses to reach target
	switchCount = 0
	for _, line := range lines {
		r := []rune(line)
		startTarget := slices.Index(r, '[')
		endTarget := slices.Index(r, ']')
		var target Vector = 0
		for i := startTarget + 1; i < endTarget; i++ {
			if r[i] == '#' {
				target = target | (1 << (i - startTarget - 1))
			}
		}
		startJolts := slices.Index(r, '{')
		switches := parseSwitches(line[endTarget+2 : startJolts+2])
		fmt.Println("targets:", target)
		fmt.Println("switches:", switches)
		var currentPos Vector = 0
		minDist := -1
		for target.Distance(currentPos) > 0 {
			//fmt.Println("currentPos:", currentPos, "target:", target, "distance:", target.Distance(currentPos))
			for _, sw := range switches {
				newPos := currentPos ^ sw // toggle switch using XOR
				//fmt.Println("currentPos:", currentPos, "trying switch:", sw, "newPos:", newPos)
				d := target.Distance(newPos)
				if minDist == -1 || d < minDist {
					minDist = d
					currentPos = newPos
					switchCount++ // count this switch press
				}
			}
		}
	}
	return
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
		var v Vector = 0
		// First number is always present
		idx1, _ := strconv.Atoi(match[1])
		v = v | (1 << idx1)
		// Check if there's a second number (pair)
		if match[2] != "" {
			idx2, _ := strconv.Atoi(match[2])
			v = v | (1 << idx2)
		}
		result = append(result, v)
	}

	return result
}
