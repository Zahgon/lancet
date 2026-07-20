package maputil

import (
	"sync"
)

const defaultShardCount = 32

type ConcurrentMap[K comparable, V any] struct {
	shardCount uint64
	locks      []sync.RWMutex
	maps       []map[K]V
}

func NewConcurrentMap[K comparable, V any](shardCount int) *ConcurrentMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (cm *ConcurrentMap[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (cm *ConcurrentMap[K, V]) GetOrSet(key K, value V) (actual V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (cm *ConcurrentMap[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

func (cm *ConcurrentMap[K, V]) GetAndDelete(key K) (actual V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (cm *ConcurrentMap[K, V]) Has(key K) bool { _ = "STUB: not implemented"; return false }

func (cm *ConcurrentMap[K, V]) Range(iterator func(key K, value V) bool) {
	_ = "STUB: not implemented"
	return
}

func (cm *ConcurrentMap[K, V]) getShard(key K) uint64 { _ = "STUB: not implemented"; return 0 }

func fnv32(key string) uint32 { _ = "STUB: not implemented"; return 0 }
