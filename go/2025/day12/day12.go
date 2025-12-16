package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mobanhawi/aoc/util"
)

/*
Solve https://adventofcode.com/2025/day/12
*/
func Solve() {
	fmt.Println("2025/day/12 pt1", solvePt1(util.ReadLines("./2025/day12/input.txt")))
}

var regionRe = regexp.MustCompile(`([0-9]+)x([0-9]+): (.+)`)
var giftRe = regexp.MustCompile(`([0-9]+):`)

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) int {
	count := 0
	maybeCount := 0 // regions that require rotation to fit
	// maps giftId to area
	gifts := map[int]int{}
	for lc, line := range lines {
		if line == "" {
			continue
		}
		giftMatch := giftRe.FindStringSubmatch(line)
		regionMatch := regionRe.FindStringSubmatch(line)
		if giftMatch != nil && regionMatch == nil {
			id, _ := strconv.Atoi(giftMatch[1])
			area := 0
			for offset := range 3 {
				area += strings.Count(lines[lc+offset+1], "#")
			}
			gifts[id] = area
			continue
		}

		if giftMatch == nil && regionMatch == nil {
			continue // not a valid line rest of shape we already processed before
		}

		if len(regionMatch) != 4 {
			panic(fmt.Sprintf("Invalid shape in line %d %v", lc, line))
		}
		l, _ := strconv.Atoi(regionMatch[1])
		w, _ := strconv.Atoi(regionMatch[2])
		regionArea := l * w
		numberBoxes := 0
		boxArea := 0
		for i, nString := range strings.Split(regionMatch[3], " ") {
			n, _ := strconv.Atoi(nString)
			numberBoxes += n
			boxArea += gifts[i] * n
		}
		// if the box can first without any rotations
		if regionArea >= numberBoxes*9 { // definitely fits
			count++
		} else if regionArea >= boxArea { // all boxes are actually less than 9 area
			// this is the maybe box that would require rotation
			maybeCount++
			fmt.Println("potential boxes", (regionArea-boxArea)/7, "regionArea", regionArea, "n", numberBoxes, "area", boxArea)
		} // otherwise it definitely does not fit
	}
	if maybeCount > 0 {
		panic("have maybe boxes, need to implement rotation logic")
	}
	//fmt.Println("count", count, "maybeCount", maybeCount)
	return count
}
