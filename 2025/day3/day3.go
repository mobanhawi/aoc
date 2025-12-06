package day3

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Solve https://adventofcode.com/2025/day/3
*/
func Solve() {
	input, err := os.ReadFile("2025/day3/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println("2025/day/3 part1", solvePt1(lines))
	fmt.Println("2025/day/3 part2", solvePt2(lines))
}

type Node struct {
	Level int
	Index int
}

// solvePt2 solves part 2 of the puzzle
// It calculates the total voltage based on the 12 levels in each line
func solvePt2(lines []string) int {
	joltage := 0
	for _, line := range lines {
		levels := make([]int, len(line))
		for i, char := range line {
			l, _ := strconv.Atoi(string(char))
			levels[i] = l
		}
		nLevels := len(levels)
		desiredLevels := 12
		// drop the n lowest levels
		for _ = range nLevels - desiredLevels {
			minIndex := 0
			for i, level := range levels {
				// if we reached the end drop the least significant digit
				if i == len(levels)-1 {
					minIndex = i
					break
				}
				// we are travesing from most significant to least significant digit
				// drop the first level that is lower than the next one
				if level < levels[i+1] {
					minIndex = i
					break
				}
			}
			// remove level at minIndex
			levels = append(levels[:minIndex], levels[minIndex+1:]...)
			// fmt.Println("removed ", minLevel, " levels now ", levels)
		}
		levelJoltage := 0
		for i := range desiredLevels {
			levelJoltage += levels[i] * int(math.Pow(10, float64((desiredLevels-1-i))))
		}
		joltage += levelJoltage
		// fmt.Println("line ", line, "total joltage ", joltage, " joltage ", levelJoltage)
	}
	return joltage
}

// solvePt1 solves part 1 of the puzzle
// It calculates the total voltage based on the two highest levels in each line
func solvePt1(lines []string) int {
	voltage := 0
	for _, line := range lines {
		levels := make([]Node, len(line))
		for i, char := range line {
			l, _ := strconv.Atoi(string(char))
			levels[i] = Node{Level: l, Index: i}
		}
		levelsWithoutLast := levels[:len(line)-1]
		element1 := slices.MaxFunc(levelsWithoutLast, func(a, b Node) int {
			return a.Level - b.Level
		})
		levelsAfterFirst := levels[element1.Index+1:]
		element2 := slices.MaxFunc(levelsAfterFirst, func(a, b Node) int {
			return a.Level - b.Level
		})
		voltage += element1.Level*10 + element2.Level
		//fmt.Println(line, "->", element1.Level, element2.Level, "->", voltage)
	}
	return voltage
}
