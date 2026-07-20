package datastructure

var defaultMapCapacity uint64 = 1 << 10

type mapNode struct {
	key   any
	value any
	next  *mapNode
}

type HashMap struct {
	capacity uint64
	size     uint64
	table    []*mapNode
}

func NewHashMap() *HashMap { _ = "STUB: not implemented"; return nil }

func NewHashMapWithCapacity(size, capacity uint64) *HashMap { _ = "STUB: not implemented"; return nil }

func (hm *HashMap) Get(key any) any { _ = "STUB: not implemented"; return *new(any) }

func (hm *HashMap) GetOrDefault(key any, defaultValue any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (hm *HashMap) Put(key any, value any) { _ = "STUB: not implemented"; return }

func (hm *HashMap) putValue(hash uint64, key, value any) { _ = "STUB: not implemented"; return }

func (hm *HashMap) Delete(key any) { _ = "STUB: not implemented"; return }

func (hm *HashMap) Contains(key any) bool { _ = "STUB: not implemented"; return false }

func (hm *HashMap) Iterate(iteratee func(key, value any)) { _ = "STUB: not implemented"; return }

func (hm *HashMap) FilterByValue(perdicate func(value any) bool) *HashMap {
	_ = "STUB: not implemented"
	return nil
}

func (hm *HashMap) Keys() []any { _ = "STUB: not implemented"; return nil }

func (hm *HashMap) Values() []any { _ = "STUB: not implemented"; return nil }

func (hm *HashMap) resize() { _ = "STUB: not implemented"; return }

func (hm *HashMap) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (hm *HashMap) hash(key any) uint64 { _ = "STUB: not implemented"; return 0 }

func newMapNode(key, value any) *mapNode { _ = "STUB: not implemented"; return nil }

func newMapNodeWithNext(key, value any, next *mapNode) *mapNode {
	_ = "STUB: not implemented"
	return nil
}
