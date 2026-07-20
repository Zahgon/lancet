package slice

import (
	"reflect"
	"sync"

	"golang.org/x/exp/constraints"
)

var (
	memoryHashMap     = make(map[string]map[any]int)
	memoryHashCounter = make(map[string]int)
	muForMemoryHash   sync.RWMutex
)

func Contain[T comparable](slice []T, target T) bool { _ = "STUB: not implemented"; return false }

func ContainBy[T any](slice []T, predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func ContainSubSlice[T comparable](slice, subSlice []T) bool {
	_ = "STUB: not implemented"
	return false
}

func ContainAny[T comparable](slice []T, targets []T) bool { _ = "STUB: not implemented"; return false }

func Chunk[T any](slice []T, size int) [][]T { _ = "STUB: not implemented"; return nil }

func Compact[T comparable](slice []T) []T { _ = "STUB: not implemented"; return nil }

func Concat[T any](slices ...[]T) []T { _ = "STUB: not implemented"; return nil }

func Difference[T comparable](slice, comparedSlice []T) []T { _ = "STUB: not implemented"; return nil }

func DifferenceBy[T comparable](slice []T, comparedSlice []T, iteratee func(index int, item T) T) []T {
	_ = "STUB: not implemented"
	return nil
}

func DifferenceWith[T any](slice []T, comparedSlice []T, comparator func(item1, item2 T) bool) []T {
	_ = "STUB: not implemented"
	return nil
}

func Equal[T comparable](slice1, slice2 []T) bool { _ = "STUB: not implemented"; return false }

func EqualWith[T, U any](slice1 []T, slice2 []U, comparator func(T, U) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualUnordered[T comparable](slice1, slice2 []T) bool { _ = "STUB: not implemented"; return false }

func Every[T any](slice []T, predicate func(index int, item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func None[T any](slice []T, predicate func(index int, item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func Some[T any](slice []T, predicate func(index int, item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	_ = "STUB: not implemented"
	return nil
}

func Count[T comparable](slice []T, item T) int { _ = "STUB: not implemented"; return 0 }

func CountBy[T any](slice []T, predicate func(index int, item T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func GroupBy[T any](slice []T, groupFn func(index int, item T) bool) ([]T, []T) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GroupWith[T any, U comparable](slice []T, iteratee func(item T) U) map[U][]T {
	_ = "STUB: not implemented"
	return nil
}

func Find[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func FindLast[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func FindBy[T any](slice []T, predicate func(index int, item T) bool) (v T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func FindLastBy[T any](slice []T, predicate func(index int, item T) bool) (v T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func Flatten(slice any) any { _ = "STUB: not implemented"; return *new(any) }

func FlattenDeep(slice any) any { _ = "STUB: not implemented"; return *new(any) }

func flattenRecursive(value reflect.Value, result reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func ForEach[T any](slice []T, iteratee func(index int, item T)) { _ = "STUB: not implemented"; return }

func ForEachWithBreak[T any](slice []T, iteratee func(index int, item T) bool) {
	_ = "STUB: not implemented"
	return
}

func Map[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	_ = "STUB: not implemented"
	return nil
}

func FilterMap[T any, U any](slice []T, iteratee func(index int, item T) (U, bool)) []U {
	_ = "STUB: not implemented"
	return nil
}

func FlatMap[T any, U any](slice []T, iteratee func(index int, item T) []U) []U {
	_ = "STUB: not implemented"
	return nil
}

func Reduce[T any](slice []T, iteratee func(index int, item1, item2 T) T, initial T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func ReduceBy[T any, U any](slice []T, initial U, reducer func(index int, item T, agg U) U) U {
	_ = "STUB: not implemented"
	return *new(U)
}

func ReduceRight[T any, U any](slice []T, initial U, reducer func(index int, item T, agg U) U) U {
	_ = "STUB: not implemented"
	return *new(U)
}

func Replace[T comparable](slice []T, old T, new T, n int) []T {
	_ = "STUB: not implemented"
	return nil
}

func ReplaceAll[T comparable](slice []T, old T, new T) []T { _ = "STUB: not implemented"; return nil }

func Repeat[T any](item T, n int) []T { _ = "STUB: not implemented"; return nil }

func InterfaceSlice(slice any) []any { _ = "STUB: not implemented"; return nil }

func StringSlice(slice any) []string { _ = "STUB: not implemented"; return nil }

func IntSlice(slice any) []int { _ = "STUB: not implemented"; return nil }

func DeleteAt[T any](slice []T, index int) []T { _ = "STUB: not implemented"; return nil }

func zeroValue[T any]() T { _ = "STUB: not implemented"; return *new(T) }

func DeleteRange[T any](slice []T, start, end int) []T { _ = "STUB: not implemented"; return nil }

func Drop[T any](slice []T, n int) []T { _ = "STUB: not implemented"; return nil }

func DropRight[T any](slice []T, n int) []T { _ = "STUB: not implemented"; return nil }

func DropWhile[T any](slice []T, predicate func(item T) bool) []T {
	_ = "STUB: not implemented"
	return nil
}

func DropRightWhile[T any](slice []T, predicate func(item T) bool) []T {
	_ = "STUB: not implemented"
	return nil
}

func InsertAt[T any](slice []T, index int, value any) []T { _ = "STUB: not implemented"; return nil }

func UpdateAt[T any](slice []T, index int, value T) []T { _ = "STUB: not implemented"; return nil }

func Unique[T comparable](slice []T) []T { _ = "STUB: not implemented"; return nil }

func UniqueBy[T any, U comparable](slice []T, iteratee func(item T) U) []T {
	_ = "STUB: not implemented"
	return nil
}

func UniqueByComparator[T comparable](slice []T, comparator func(item T, other T) bool) []T {
	_ = "STUB: not implemented"
	return nil
}

func UniqueByField[T any](slice []T, field string) ([]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getField[T any](item T, field string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Union[T comparable](slices ...[]T) []T { _ = "STUB: not implemented"; return nil }

func UnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T {
	_ = "STUB: not implemented"
	return nil
}

func Merge[T any](slices ...[]T) []T { _ = "STUB: not implemented"; return nil }

func Intersection[T comparable](slices ...[]T) []T { _ = "STUB: not implemented"; return nil }

func SymmetricDifference[T comparable](slices ...[]T) []T { _ = "STUB: not implemented"; return nil }

func Reverse[T any](slice []T) { _ = "STUB: not implemented"; return }

func ReverseCopy[T any](slice []T) []T { _ = "STUB: not implemented"; return nil }

func Shuffle[T any](slice []T) []T { _ = "STUB: not implemented"; return nil }

func ShuffleCopy[T any](slice []T) []T { _ = "STUB: not implemented"; return nil }

func IsAscending[T constraints.Ordered](slice []T) bool { _ = "STUB: not implemented"; return false }

func IsDescending[T constraints.Ordered](slice []T) bool { _ = "STUB: not implemented"; return false }

func IsSorted[T constraints.Ordered](slice []T) bool { _ = "STUB: not implemented"; return false }

func IsSortedByKey[T any, K constraints.Ordered](slice []T, iteratee func(item T) K) bool {
	_ = "STUB: not implemented"
	return false
}

func Sort[T constraints.Ordered](slice []T, sortOrder ...string) { _ = "STUB: not implemented"; return }

func SortBy[T any](slice []T, less func(a, b T) bool) { _ = "STUB: not implemented"; return }

func SortByField[T any](slice []T, field string, sortType ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func Without[T comparable](slice []T, items ...T) []T { _ = "STUB: not implemented"; return nil }

func IndexOf[T comparable](arr []T, val T) int { _ = "STUB: not implemented"; return 0 }

func LastIndexOf[T comparable](slice []T, item T) int { _ = "STUB: not implemented"; return 0 }

func ToSlicePointer[T any](items ...T) []*T { _ = "STUB: not implemented"; return nil }

func ToSlice[T any](items ...T) []T { _ = "STUB: not implemented"; return nil }

func AppendIfAbsent[T comparable](slice []T, item T) []T { _ = "STUB: not implemented"; return nil }

func SetToDefaultIf[T any](slice []T, predicate func(T) bool) ([]T, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func KeyBy[T any, U comparable](slice []T, iteratee func(item T) U) map[U]T {
	_ = "STUB: not implemented"
	return nil
}

func Join[T any](slice []T, separator string) string { _ = "STUB: not implemented"; return "" }

func Partition[T any](slice []T, predicates ...func(item T) bool) [][]T {
	_ = "STUB: not implemented"
	return nil
}

func Break[T any](values []T, predicate func(T) bool) ([]T, []T) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Random[T any](slice []T) (val T, idx int) { _ = "STUB: not implemented"; return *new(T), 0 }

func RightPadding[T any](slice []T, paddingValue T, paddingLength int) []T {
	_ = "STUB: not implemented"
	return nil
}

func LeftPadding[T any](slice []T, paddingValue T, paddingLength int) []T {
	_ = "STUB: not implemented"
	return nil
}

func Frequency[T comparable](slice []T) map[T]int { _ = "STUB: not implemented"; return nil }

func JoinFunc[T any](slice []T, sep string, transform func(T) T) string {
	_ = "STUB: not implemented"
	return ""
}

func ConcatBy[T any](slice []T, sep T, connector func(T, T) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}
