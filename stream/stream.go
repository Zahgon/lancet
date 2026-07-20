package stream

import (
	"golang.org/x/exp/constraints"
)

type Stream[T any] struct {
	source []T
}

func Of[T any](elems ...T) Stream[T] { _ = "STUB: not implemented"; return nil }

func Generate[T any](generator func() func() (item T, ok bool)) Stream[T] {
	_ = "STUB: not implemented"
	return nil
}

func FromSlice[T any](source []T) Stream[T] { _ = "STUB: not implemented"; return nil }

func FromChannel[T any](source <-chan T) Stream[T] { _ = "STUB: not implemented"; return nil }

func FromRange[T constraints.Integer | constraints.Float](start, end, step T) Stream[T] {
	_ = "STUB: not implemented"
	return nil
}

func Concat[T any](a, b Stream[T]) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Distinct() Stream[T] { _ = "STUB: not implemented"; return nil }

func hashKey(data any) string { _ = "STUB: not implemented"; return "" }

func (s Stream[T]) Filter(predicate func(item T) bool) Stream[T] {
	_ = "STUB: not implemented"
	return nil
}

func (s Stream[T]) Map(mapper func(item T) T) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Peek(consumer func(item T)) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Skip(n int) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Limit(maxSize int) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) AllMatch(predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Stream[T]) AnyMatch(predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Stream[T]) NoneMatch(predicate func(item T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Stream[T]) ForEach(action func(item T)) { _ = "STUB: not implemented"; return }

func (s Stream[T]) Reduce(initial T, accumulator func(a, b T) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func (s Stream[T]) Count() int { _ = "STUB: not implemented"; return 0 }

func (s Stream[T]) FindFirst() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s Stream[T]) FindLast() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s Stream[T]) Reverse() Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Range(start, end int) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Sorted(less func(a, b T) bool) Stream[T] { _ = "STUB: not implemented"; return nil }

func (s Stream[T]) Max(less func(a, b T) bool) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (s Stream[T]) Min(less func(a, b T) bool) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (s Stream[T]) IndexOf(target T, equal func(a, b T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func (s Stream[T]) LastIndexOf(target T, equal func(a, b T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func (s Stream[T]) ToSlice() []T { _ = "STUB: not implemented"; return nil }

func ToMap[T any, K comparable, V any](s Stream[T], mapper func(item T) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}
