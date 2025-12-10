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
