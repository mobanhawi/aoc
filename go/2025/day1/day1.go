package day1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mobanhawi/aoc/util"
)

/*

Solve https://adventofcode.com/2025/day/1

Counting dial positions at 0
Dial numbered 0-99, starts at 50
Instructions like "L68" (left 68) or "R30" (right 30)
Count how many times the dial lands exactly on 0

You remember from the training seminar that "method 0x434C49434B" means
you're actually supposed to count the number of times any click causes the dial to point at 0,
regardless of whether it happens during a rotation or at the end of one.




*/

func Solve() {
	fmt.Println("2025/day1", solve(util.ReadLines("2025/day1/input.txt"), 50))
}

func solve(sequence []string, position int) int {
	count := 0

	for _, instruction := range sequence {
		if len(instruction) < 2 {
			log.Fatalf("invalid instruction: %s", instruction)
		}

		direction := instruction[0]
		clicks, _ := strconv.Atoi(instruction[1:])
		// Count each individual click that passes through 0
		for i := 0; i < clicks; i++ {
			if direction == 'L' {
				position--
				if position < 0 {
					position = 99
				}
			} else if direction == 'R' {
				position++
				if position > 99 {
					position = 0
				}
			}

			if position == 0 {
				count++
			}
		}
	}

	return count
}
