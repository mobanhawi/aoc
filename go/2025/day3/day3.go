package day3

import (
	"fmt"
	"math"

	"github.com/mobanhawi/aoc/util"
)

/*
Solve https://adventofcode.com/2025/day/3
*/
func Solve() {
	lines := util.ReadLines("2025/day3/input.txt")
	fmt.Println("2025/day/3 part1", solvePt1(lines))
	fmt.Println("2025/day/3 part2", solvePt2(lines))
}

// solveWithNumBatteries finds the maximum joltage by keeping the n highest batteries levels from each bank
// and dropping the rest
// It returns the total joltage across all banks
func solveWithNumBatteries(lines []string, numBatteries int) int {
	joltage := 0
	for _, line := range lines {
		banks := make([]int, len(line))
		for i, char := range line {
			l := int(char - '0')
			banks[i] = l
		}
		// drop the n lowest levels
		for _ = range len(banks) - numBatteries {
			minIndex := 0
			for i, level := range banks {
				// if we reached the end drop the least significant digit
				if i == len(banks)-1 {
					minIndex = i
					break
				}
				// we are travesing from most significant to least significant digit
				// drop the first level that is lower than the next one
				if level < banks[i+1] {
					minIndex = i
					break
				}
			}
			// remove battery at minIndex
			banks = append(banks[:minIndex], banks[minIndex+1:]...)
			// fmt.Println("removed ", minLevel, " levels now ", levels)
		}
		bankJoltage := 0
		for i := range numBatteries {
			// add level joltage for this bank from most significant to least significant digit
			bankJoltage += banks[i] * int(math.Pow(10, float64(numBatteries-1-i)))
		}
		joltage += bankJoltage
		// fmt.Println("line ", line, "total joltage ", joltage, " joltage ", levelJoltage)
	}
	return joltage
}

// solvePt2 solves part 2 of the puzzle
// It calculates the total voltage based on the 12 levels in each line
func solvePt2(lines []string) int {
	return solveWithNumBatteries(lines, 12)
}

// solvePt1 solves part 1 of the puzzle
// It calculates the total voltage based on the two highest levels in each line
func solvePt1(lines []string) int {
	return solveWithNumBatteries(lines, 2)
}
