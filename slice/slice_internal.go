package slice

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

type resultChunk[T comparable] struct {
	index int
	data  []T
}

func sliceValue(slice any) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func sliceElemType(reflectType reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func quickSort[T constraints.Ordered](slice []T, lowIndex, highIndex int, order string) {
	_ = "STUB: not implemented"
	return
}

func partitionOrderedSlice[T constraints.Ordered](slice []T, lowIndex, highIndex int, order string) int {
	_ = "STUB: not implemented"
	return 0
}

func quickSortBy[T any](slice []T, lowIndex, highIndex int, less func(a, b T) bool) {
	_ = "STUB: not implemented"
	return
}

func partitionAnySlice[T any](slice []T, lowIndex, highIndex int, less func(a, b T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func swap[T any](slice []T, i, j int) { _ = "STUB: not implemented"; return }

func repeat[S ~[]E, E any](x S, count int) S { _ = "STUB: not implemented"; return *new(S) }

func concat[S ~[]E, E any](slices ...S) S { _ = "STUB: not implemented"; return *new(S) }

func grow[S ~[]E, E any](s S, n int) S { _ = "STUB: not implemented"; return *new(S) }
