package datastructure

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] { _ = "STUB: not implemented"; return nil }

func FromSlice[T comparable](items []T) Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) Add(items ...T) { _ = "STUB: not implemented"; return }

func (s Set[T]) AddIfNotExist(item T) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) AddIfNotExistBy(item T, checker func(element T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Set[T]) Contain(item T) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) ContainAll(other Set[T]) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) Clone() Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) Delete(items ...T) { _ = "STUB: not implemented"; return }

func (s Set[T]) Equal(other Set[T]) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) Iterate(fn func(item T)) { _ = "STUB: not implemented"; return }

func (s Set[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (s Set[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (s Set[T]) Union(other Set[T]) Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) Intersection(other Set[T]) Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) Minus(comparedSet Set[T]) Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) EachWithBreak(iteratee func(item T) bool) { _ = "STUB: not implemented"; return }

func (s Set[T]) Pop() (v T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (s Set[T]) ToSlice() []T { _ = "STUB: not implemented"; return nil }

func (s Set[T]) ToSortedSlice(less func(v1, v2 T) bool) []T { _ = "STUB: not implemented"; return nil }
