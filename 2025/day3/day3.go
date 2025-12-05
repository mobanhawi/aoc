package day3

import (
	"fmt"
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
	fmt.Println("2025/day/3", solve(lines))
}

type Node struct {
	Level int
	Index int
}

func solve(lines []string) int {
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
