package algorithm

import "github.com/duke-git/lancet/v2/constraints"

func LinearSearch[T any](slice []T, target T, equal func(a, b T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func BinarySearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator constraints.Comparator) int {
	_ = "STUB: not implemented"
	return 0
}

func BinaryIterativeSearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator constraints.Comparator) int {
	_ = "STUB: not implemented"
	return 0
}
