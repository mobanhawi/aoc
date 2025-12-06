package day4

import (
	"fmt"
	"os"
	"strings"
)

/*
Solve https://adventofcode.com/2025/day/4
*/
func Solve() {
	input, err := os.ReadFile("2025/day4/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println("2025/day/4", solve(lines))
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
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if board[x][y] && adjancentOccupancy(board, x, y) < 4 {
				access++
			}
		}
	}
	return access
}

// if there are fewer than four rolls of paper in the eight adjacent positions
func adjancentOccupancy(board Board, x, y int) int {
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
