package datastructure

import (
	"github.com/duke-git/lancet/v2/constraints"
)

type MaxHeap[T any] struct {
	data       []T
	comparator constraints.Comparator
}

func NewMaxHeap[T any](comparator constraints.Comparator) *MaxHeap[T] {
	_ = "STUB: not implemented"
	return nil
}

func BuildMaxHeap[T any](data []T, comparator constraints.Comparator) *MaxHeap[T] {
	_ = "STUB: not implemented"
	return nil
}

func (h *MaxHeap[T]) Push(value T) { _ = "STUB: not implemented"; return }

func (h *MaxHeap[T]) heapifyUp(i int) { _ = "STUB: not implemented"; return }

func (h *MaxHeap[T]) Pop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (h *MaxHeap[T]) heapifyDown(i int) { _ = "STUB: not implemented"; return }

func (h *MaxHeap[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (h *MaxHeap[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (h *MaxHeap[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (h *MaxHeap[T]) PrintStructure() { _ = "STUB: not implemented"; return }

func parentIndex(i int) int { _ = "STUB: not implemented"; return 0 }

func leftChildIndex(i int) int { _ = "STUB: not implemented"; return 0 }

func rightChildIndex(i int) int { _ = "STUB: not implemented"; return 0 }

func (h *MaxHeap[T]) swap(i, j int) { _ = "STUB: not implemented"; return }

func powerTwo(n int) int { _ = "STUB: not implemented"; return 0 }
