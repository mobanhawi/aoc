package day4

import (
	"fmt"

	"github.com/mobanhawi/aoc/2025/util"
)

/*
Solve https://adventofcode.com/2025/day/4
*/
func Solve() {
	fmt.Println("2025/day/4", solve(util.ReadLines("2025/day4/input.txt")))
}

type Board map[int]map[int]bool

func solve(lines []string) int {
	// Initialize board
	board := make(Board)
	access := 0
	for y, line := range lines {
		for x, char := range line {
			if board[x] == nil {
				board[x] = make(map[int]bool)
			}
			board[x][y] = char == '@'
		}
	}
	lenX := len(lines[0])
	lenY := len(lines)
	// Implement logic to solve the puzzle here
	for {
		removed := countAndRemove(board, lenX, lenY)
		if removed == 0 {
			break
		}
		access += removed
	}

	return access
}

// countAndRemove counts and removes rolls of paper that have fewer than four adjacent rolls
func countAndRemove(board Board, lenX, lenY int) int {
	count := 0
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if board[x][y] && adjacentOccupancy(board, x, y) < 4 {
				count++
				board[x][y] = false // remove roll of paper
			}
		}
	}
	return count
}

// adjacentOccupancy counts the number of adjacent occupied rolls of paper
func adjacentOccupancy(board Board, x, y int) int {
	count := 0
	// move in 8 directions
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if board[nx] != nil && board[nx][ny] {
				count++
			}
		}
	}
	return count
}
