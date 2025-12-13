package util

import (
	"os"
	"strings"
)

func ReadLines(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(input)), "\n")
}

func ReadLinesNoTrim(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(input), "\n")
}
