package datastructure

import (
	"github.com/duke-git/lancet/v2/constraints"
)

type PriorityQueue[T any] struct {
	items      []T
	size       int
	comparator constraints.Comparator
}

func NewPriorityQueue[T any](capacity int, comparator constraints.Comparator) *PriorityQueue[T] {
	_ = "STUB: not implemented"
	return nil
}

func (q *PriorityQueue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *PriorityQueue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (q *PriorityQueue[T]) IsFull() bool { _ = "STUB: not implemented"; return false }

func (q *PriorityQueue[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (q *PriorityQueue[T]) Enqueue(val T) error { _ = "STUB: not implemented"; return nil }

func (q *PriorityQueue[T]) Dequeue() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *PriorityQueue[T]) swim(index int) { _ = "STUB: not implemented"; return }

func (q *PriorityQueue[T]) sink(index int) { _ = "STUB: not implemented"; return }

func (q *PriorityQueue[T]) swap(i, j int) { _ = "STUB: not implemented"; return }
