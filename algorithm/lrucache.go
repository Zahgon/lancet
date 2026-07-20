package algorithm

type lruNode[K comparable, V any] struct {
	key   K
	value V
	pre   *lruNode[K, V]
	next  *lruNode[K, V]
}

func newLruNode[K comparable, V any](key K, value V) *lruNode[K, V] {
	_ = "STUB: not implemented"
	return nil
}

type LRUCache[K comparable, V any] struct {
	cache    map[K]*lruNode[K, V]
	head     *lruNode[K, V]
	tail     *lruNode[K, V]
	capacity int
	length   int
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (l *LRUCache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (l *LRUCache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (l *LRUCache[K, V]) Delete(key K) bool { _ = "STUB: not implemented"; return false }

func (l *LRUCache[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *LRUCache[K, V]) addNode(node *lruNode[K, V]) { _ = "STUB: not implemented"; return }

func (l *LRUCache[K, V]) deleteNode(node *lruNode[K, V]) K {
	_ = "STUB: not implemented"
	return *new(K)
}

func (l *LRUCache[K, V]) moveToTail(node *lruNode[K, V]) { _ = "STUB: not implemented"; return }
