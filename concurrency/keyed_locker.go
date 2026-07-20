package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type KeyedLocker[K comparable] struct {
	locks sync.Map
	ttl   time.Duration
}

type lockEntry struct {
	mu    sync.Mutex
	ref   int32
	timer atomic.Pointer[time.Timer]
}

func NewKeyedLocker[K comparable](ttl time.Duration) *KeyedLocker[K] {
	_ = "STUB: not implemented"
	return nil
}

func (l *KeyedLocker[K]) Do(ctx context.Context, key K, fn func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *KeyedLocker[K]) acquire(key K) *lockEntry { _ = "STUB: not implemented"; return nil }

func (l *KeyedLocker[K]) release(key K, entry *lockEntry, rawKey K) {
	_ = "STUB: not implemented"
	return
}

type RWKeyedLocker[K comparable] struct {
	locks sync.Map
	ttl   time.Duration
}

type rwLockEntry struct {
	mu    sync.RWMutex
	ref   int32
	timer atomic.Pointer[time.Timer]
}

func NewRWKeyedLocker[K comparable](ttl time.Duration) *RWKeyedLocker[K] {
	_ = "STUB: not implemented"
	return nil
}

func (l *RWKeyedLocker[K]) RLock(ctx context.Context, key K, fn func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *RWKeyedLocker[K]) Lock(ctx context.Context, key K, fn func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *RWKeyedLocker[K]) acquire(key K) *rwLockEntry { _ = "STUB: not implemented"; return nil }

func (l *RWKeyedLocker[K]) release(entry *rwLockEntry, rawKey K) { _ = "STUB: not implemented"; return }

type TryKeyedLocker[K comparable] struct {
	mu    sync.Mutex
	locks map[K]*casMutex
}

func NewTryKeyedLocker[K comparable]() *TryKeyedLocker[K] { _ = "STUB: not implemented"; return nil }

func (l *TryKeyedLocker[K]) TryLock(key K) bool { _ = "STUB: not implemented"; return false }

func (l *TryKeyedLocker[K]) Unlock(key K) { _ = "STUB: not implemented"; return }

type casMutex struct {
	lock int32
}

func (m *casMutex) TryLock() bool { _ = "STUB: not implemented"; return false }

func (m *casMutex) Unlock() { _ = "STUB: not implemented"; return }
