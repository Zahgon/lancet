package enum

import (
	"sync"
)

type Enum[T comparable] interface {
	Value() T
	String() string
	Name() string
	Valid(checker ...func(T) bool) bool
}

type Item[T comparable] struct {
	value T
	name  string
}

func NewItem[T comparable](value T, name string) *Item[T] { _ = "STUB: not implemented"; return nil }

type Pair[T comparable] struct {
	Value T
	Name  string
}

func NewItemsFromPairs[T comparable](pairs ...Pair[T]) []*Item[T] {
	_ = "STUB: not implemented"
	return nil
}

func (it *Item[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (it *Item[T]) Name() string { _ = "STUB: not implemented"; return "" }

func (it *Item[T]) String() string { _ = "STUB: not implemented"; return "" }

func (it *Item[T]) Valid(checker ...func(T) bool) bool { _ = "STUB: not implemented"; return false }

func (it *Item[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (it *Item[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type Registry[T comparable] struct {
	mu     sync.RWMutex
	values map[T]*Item[T]
	names  map[string]*Item[T]
	items  []*Item[T]
}

func NewRegistry[T comparable](items ...*Item[T]) *Registry[T] {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry[T]) Add(items ...*Item[T]) { _ = "STUB: not implemented"; return }

func (r *Registry[T]) Remove(value T) bool { _ = "STUB: not implemented"; return false }

func (r *Registry[T]) Update(value T, newName string) bool { _ = "STUB: not implemented"; return false }

func (r *Registry[T]) GetByValue(value T) (*Item[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *Registry[T]) GetByName(name string) (*Item[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *Registry[T]) Items() []*Item[T] { _ = "STUB: not implemented"; return nil }

func (r *Registry[T]) Contains(value T) bool { _ = "STUB: not implemented"; return false }

func (r *Registry[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (r *Registry[T]) Range(fn func(*Item[T]) bool) { _ = "STUB: not implemented"; return }

func (r *Registry[T]) SortedItems(less func(*Item[T], *Item[T]) bool) []*Item[T] {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry[T]) Filter(predicate func(*Item[T]) bool) []*Item[T] {
	_ = "STUB: not implemented"
	return nil
}
