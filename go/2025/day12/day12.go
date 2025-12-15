package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mobanhawi/aoc/util"
)

/*
Solve https://adventofcode.com/2025/day/10
*/
func Solve() {
	fmt.Println("2025/day/11 pt1", solvePt1(util.ReadLines("./2025/day12/input.txt")))
	fmt.Println("2025/day/11 pt2", solvePt2(util.ReadLines("./2025/day12/input.txt")))
}

var regionRe = regexp.MustCompile(`([0-9]+)x([0-9]+): (.+)`)
var giftRe = regexp.MustCompile(`([0-9]+):`)

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) int {
	count := 0
	// maps giftId to area
	//gifts := map[int]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		//if matches := giftRe.FindStringSubmatch(line); matches != nil {
		//	id, _ := strconv.Atoi(matches[1])
		//	area := 0
		//	for offset := range 3 {
		//		area += strings.Count(lines[c+offset+1], "#")
		//	}
		//	gifts[id] = area
		//	fmt.Println(id, area)
		//	continue
		//}

		matches := regionRe.FindStringSubmatch(line)
		if len(matches) != 4 {
			fmt.Println("Invalid shape in line", line)
			continue
		}
		l, _ := strconv.Atoi(matches[1])
		w, _ := strconv.Atoi(matches[2])
		regionArea := l * w
		numberBoxes := 0
		for _, nString := range strings.Split(matches[3], " ") {
			n, _ := strconv.Atoi(nString)
			numberBoxes += n
		}
		// if the box can first without any rotations
		if regionArea >= numberBoxes*9 { // define
			count++
		} else if regionArea >= numberBoxes*7 { // all boxes are actually 7 #s
			// this is the maybe box that would require rotation
			fmt.Println("potential boxes", (regionArea-numberBoxes*7)/7, "regionArea", regionArea, "n", numberBoxes, "area", numberBoxes*7)
		} else {
			fmt.Println("not valid boxes", (regionArea-numberBoxes*7)/7, "regionArea", regionArea, "n", numberBoxes, "area", numberBoxes*7)
		}
	}
	return count
}

// solvePt1 solves part 2 of the puzzle
func solvePt2(lines []string) int {
	return 0
}
