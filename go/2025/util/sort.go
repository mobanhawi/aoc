package util

import "golang.org/x/exp/constraints"

// Sort returns the two values in ascending order.
func Sort[T constraints.Ordered](a, b T) (T, T) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func FindMinIndex[T constraints.Ordered](values []T) int {
	if len(values) == 0 {
		return -1 // Or handle error for empty slice
	}
	minIndex := 0
	for i, v := range values {
		if v < values[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
