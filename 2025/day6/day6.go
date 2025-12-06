package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mobanhawi/aoc/2025/util"
)

/*
Solve https://adventofcode.com/2025/day/6
*/
func Solve() {
	fmt.Println("2025/day/6 pt1", solvePt1(util.ReadLines("./2025/day6/input1.txt")))
	fmt.Println("2025/day/6 pt2", solvePt2(util.ReadLines("./2025/day6/input0.txt")))
}

// solvePt2 solves part 1 of the puzzle
func solvePt2(lines []string) int {
	cols := len(lines[0])
	rows := len(lines)
	result := 0
	total := 0
	sign := ""

	for col := range cols {
		buffer := 0
		for row := range rows {
			char := lines[row][col]
			if char >= '0' && char <= '9' {
				buffer = buffer*10 + int(char-'0')
			}
			if row == rows-1 { // last row
				// apply last operator
				if char == '+' || char == '*' {
					sign = string(char)
					result = buffer
				}
				// blank col
				if buffer == 0 && char == ' ' {
					total += result
					result = 0
					sign = ""
					continue
				}
				if char == ' ' {
					switch sign {
					case "+":
						result += buffer
					case "*":
						result *= buffer
					default:
						panic("unknown operator")
					}
				}
			}
			fmt.Println("col", col, "row", row, "char", string(char), "buffer", buffer, "result", result, "total", total)

		}
	}
	return 0
}

// solvePt1 solves part 1 of the puzzle
func solvePt1(lines []string) int {
	nProblems := len(getNumbersFromLine(lines[0]))
	nNumbers := len(lines) - 1 // last char is operator
	// numbers[i][j] is number i for problem j
	numbers := make([][]int, nNumbers)
	operator := make([]string, nProblems)
	for i, line := range lines { // number from each line
		if i == nNumbers { // last line is operator
			operator = getOperatorFromLine(line)
			break
		}
		numbers[i] = make([]int, nProblems)
		for j, num := range getNumbersFromLine(line) {
			numbers[i][j] = num
		}
	}
	total := 0
	result := 0
	for j := range nProblems {
		if operator[j] == "*" {
			result = 1
			for i := range nNumbers {
				result *= numbers[i][j]
			}
		} else if operator[j] == "+" {
			result = 0
			for i := range nNumbers {
				result += numbers[i][j]
			}
		} else {
			panic("unknown operator")
		}
		total += result
	}
	return total
}

func getNumbersFromLine(line string) []int {
	parts := strings.Split(strings.TrimSpace(line), " ")
	numbers := make([]int, 0)
	for _, part := range parts {
		if part == "" {
			continue
		}
		num, _ := strconv.Atoi(part)
		numbers = append(numbers, num)
	}
	return numbers
}

func getOperatorFromLine(line string) []string {
	parts := strings.Split(strings.TrimSpace(line), " ")
	operator := make([]string, 0)
	for _, part := range parts {
		if part == "" {
			continue
		}
		operator = append(operator, part)
	}
	return operator
}
