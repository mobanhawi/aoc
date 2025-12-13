package day7

import (
	"fmt"
	"strings"

	util2 "github.com/mobanhawi/aoc/util"
)

/*
Solve https://adventofcode.com/2025/day/6
*/
func Solve() {
	fmt.Println("2025/day/7 pt1", solvePt1(util2.ReadLines("./2025/day7/input.txt"), false))
	fmt.Println("2025/day/7 pt2", solvePt2(util2.ReadLines("./2025/day7/input.txt")))
}

func solvePt2(lines []string) int {
	rows := len(lines)
	//cols := len(lines[0])
	startCol := strings.Index(lines[0], "S")

	// timelines count map
	// map col -> count of beams reaching that col
	count := map[int]int{startCol: 1}

	// process each remaining row
	for row := range rows {
		next := make(map[int]int)
		// new cols extends existing timelines
		for col, n := range count {
			if lines[row][col] == '^' {
				next[col-1] += n
				next[col+1] += n
			} else {
				next[col] += n
			}
		}
		count = next
	}

	total := 0
	for _, c := range count {
		total += c
	}
	return total
}

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string, debug bool) int {
	rows := len(lines)
	cols := len(lines[0])
	startCol := strings.Index(lines[0], "S")
	splitCount := 0
	// record beam positions per row
	beamPos := make(map[int]*util2.Set[int])
	for r := range rows {
		beamPos[r] = util2.NewSet[int]()
		if r == 0 { // starting row only
			beamPos[r].Add(startCol)
			continue
		}
		// for each beam position in previous row, check possible splits
		for _, col := range beamPos[r-1].List() {
			if lines[r][col] == '^' {
				// beam continues split
				splitCount++
				// check left split
				if col > 0 {
					beamPos[r].Add(col - 1)
				}
				// check right split
				if col < cols-1 {
					beamPos[r].Add(col + 1)
				}
			} else {
				// no split, beam continues straight
				beamPos[r].Add(col)
			}
		}
	}
	if debug {
		displayGrid(lines, beamPos)
	}
	return splitCount
}

func displayGrid(lines []string, beamPos map[int]*util2.Set[int]) {
	for r, line := range lines {
		rowChars := []rune(line)
		for _, col := range beamPos[r].List() {
			if rowChars[col] == 'S' {
				continue
			}
			rowChars[col] = '|'
		}
		fmt.Println(string(rowChars))
	}
}
