package maputil

import (
	"container/list"
	"sync"
)

type OrderedMap[K comparable, V any] struct {
	mu sync.RWMutex

	data  map[K]V
	order *list.List
	index map[K]*list.Element
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] { _ = "STUB: not implemented"; return nil }

func (om *OrderedMap[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (om *OrderedMap[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (om *OrderedMap[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

func (om *OrderedMap[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (om *OrderedMap[K, V]) Front() (struct {
	Key   K
	Value V
}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (om *OrderedMap[K, V]) Back() (struct {
	Key   K
	Value V
}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (om *OrderedMap[K, V]) Range(iteratee func(key K, value V) bool) {
	_ = "STUB: not implemented"
	return
}

func (om *OrderedMap[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (om *OrderedMap[K, V]) Values() []V { _ = "STUB: not implemented"; return nil }

func (om *OrderedMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (om *OrderedMap[K, V]) Contains(key K) bool { _ = "STUB: not implemented"; return false }

func (om *OrderedMap[K, V]) Elements() []struct {
	Key   K
	Value V
} {
	_ = "STUB: not implemented"
	return nil
}

func (om *OrderedMap[K, V]) Iter() <-chan struct {
	Key   K
	Value V
} {
	_ = "STUB: not implemented"
	return nil
}

func (om *OrderedMap[K, V]) ReverseIter() <-chan struct {
	Key   K
	Value V
} {
	_ = "STUB: not implemented"
	return nil
}

func (om *OrderedMap[K, V]) SortByKey(less func(a, b K) bool) { _ = "STUB: not implemented"; return }

func (om *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (om *OrderedMap[K, V]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func keyToString[K any](key K) (string, error) { _ = "STUB: not implemented"; return "", nil }

func stringToKey[K any](s string) (K, error) { _ = "STUB: not implemented"; return *new(K), nil }
