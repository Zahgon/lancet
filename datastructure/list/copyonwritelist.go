package datastructure

import (
	"sync"
)

type CopyOnWriteList[T any] struct {
	data []T
	lock sync.Locker
}

func NewCopyOnWriteList[T any](data []T) *CopyOnWriteList[T] { _ = "STUB: not implemented"; return nil }

func (c *CopyOnWriteList[T]) getList() []T { _ = "STUB: not implemented"; return nil }

func (c *CopyOnWriteList[T]) setList(data []T) { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *CopyOnWriteList[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (c *CopyOnWriteList[T]) Contain(e T) bool { _ = "STUB: not implemented"; return false }

func (c *CopyOnWriteList[T]) ValueOf(index int) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *CopyOnWriteList[T]) IndexOf(e T) int { _ = "STUB: not implemented"; return 0 }

func indexOf[T any](o T, e []T, start int, end int) int { _ = "STUB: not implemented"; return 0 }

func (c *CopyOnWriteList[T]) LastIndexOf(e T) int { _ = "STUB: not implemented"; return 0 }

func lastIndexOf[T any](o T, e []T, start int, end int) int { _ = "STUB: not implemented"; return 0 }

func (l *CopyOnWriteList[T]) LastIndexOfFunc(f func(T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func (l *CopyOnWriteList[T]) IndexOfFunc(f func(T) bool) int { _ = "STUB: not implemented"; return 0 }

func get[T any](o []T, index int) *T { _ = "STUB: not implemented"; return nil }

func (c *CopyOnWriteList[T]) Get(index int) *T { _ = "STUB: not implemented"; return nil }

func (c *CopyOnWriteList[T]) set(index int, e T) (oldValue *T) {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopyOnWriteList[T]) Set(index int, e T) (oldValue *T, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *CopyOnWriteList[T]) Add(e T) bool { _ = "STUB: not implemented"; return false }

func (c *CopyOnWriteList[T]) AddAll(e []T) bool { _ = "STUB: not implemented"; return false }

func (c *CopyOnWriteList[T]) AddByIndex(index int, e T) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopyOnWriteList[T]) delete(index int) *T { _ = "STUB: not implemented"; return nil }

func (c *CopyOnWriteList[T]) DeleteAt(index int) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *CopyOnWriteList[T]) DeleteBy(o T) (*T, bool) { _ = "STUB: not implemented"; return nil, false }

func (c *CopyOnWriteList[T]) DeleteRange(start int, end int) { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) DeleteIf(f func(T) bool) { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) Equal(other *[]T) bool { _ = "STUB: not implemented"; return false }

func (c *CopyOnWriteList[T]) Clear() { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) Merge(other []T) { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) ForEach(f func(T)) { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) Sort(compare func(o1 T, o2 T) bool) { _ = "STUB: not implemented"; return }

func (c *CopyOnWriteList[T]) SubList(start int, end int) (newList []T) {
	_ = "STUB: not implemented"
	return nil
}
