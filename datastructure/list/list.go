package datastructure

import (
	"github.com/duke-git/lancet/v2/iterator"
)

type List[T any] struct {
	data []T
}

func NewList[T any](data []T) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (l *List[T]) ValueOf(index int) (*T, bool) { _ = "STUB: not implemented"; return nil, false }

func (l *List[T]) IndexOf(value T) int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) LastIndexOf(value T) int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) IndexOfFunc(f func(T) bool) int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) LastIndexOfFunc(f func(T) bool) int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) Contain(value T) bool { _ = "STUB: not implemented"; return false }

func (l *List[T]) Push(value T) { _ = "STUB: not implemented"; return }

func (l *List[T]) InsertAtFirst(value T) { _ = "STUB: not implemented"; return }

func (l *List[T]) InsertAtLast(value T) { _ = "STUB: not implemented"; return }

func (l *List[T]) InsertAt(index int, value T) { _ = "STUB: not implemented"; return }

func (l *List[T]) PopFirst() (*T, bool) { _ = "STUB: not implemented"; return nil, false }

func (l *List[T]) PopLast() (*T, bool) { _ = "STUB: not implemented"; return nil, false }

func (l *List[T]) DeleteAt(index int) { _ = "STUB: not implemented"; return }

func (l *List[T]) DeleteIf(f func(T) bool) int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) UpdateAt(index int, value T) { _ = "STUB: not implemented"; return }

func (l *List[T]) Equal(other *List[T]) bool { _ = "STUB: not implemented"; return false }

func (l *List[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (l *List[T]) Clear() { _ = "STUB: not implemented"; return }

func (l *List[T]) Clone() *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) Merge(other *List[T]) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) Cap() int { _ = "STUB: not implemented"; return 0 }

func (l *List[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (l *List[T]) Reverse() { _ = "STUB: not implemented"; return }

func (l *List[T]) Unique() { _ = "STUB: not implemented"; return }

func (l *List[T]) Union(other *List[T]) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) Intersection(other *List[T]) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) Difference(other *List[T]) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) SymmetricDifference(other *List[T]) *List[T] {
	_ = "STUB: not implemented"
	return nil
}

func (l *List[T]) SubList(fromIndex, toIndex int) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) ForEach(consumer func(T)) { _ = "STUB: not implemented"; return }

func (l *List[T]) RetainAll(list *List[T]) bool { _ = "STUB: not implemented"; return false }

func (l *List[T]) DeleteAll(list *List[T]) bool { _ = "STUB: not implemented"; return false }

func (l *List[T]) batchRemove(list *List[T], complement bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *List[T]) Iterator() iterator.Iterator[T] { _ = "STUB: not implemented"; return nil }

func ListToMap[T any, K comparable, V any](list *List[T], iteratee func(T) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}
