package datastructure

type ArrayQueue[T any] struct {
	data     []T
	head     int
	tail     int
	capacity int
	size     int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] { _ = "STUB: not implemented"; return nil }

func (q *ArrayQueue[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (q *ArrayQueue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (q *ArrayQueue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *ArrayQueue[T]) IsFull() bool { _ = "STUB: not implemented"; return false }

func (q *ArrayQueue[T]) Front() T { _ = "STUB: not implemented"; return *new(T) }

func (q *ArrayQueue[T]) Back() T { _ = "STUB: not implemented"; return *new(T) }

func (q *ArrayQueue[T]) Enqueue(item T) bool { _ = "STUB: not implemented"; return false }

func (q *ArrayQueue[T]) Dequeue() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *ArrayQueue[T]) Clear() { _ = "STUB: not implemented"; return }

func (q *ArrayQueue[T]) Contain(value T) bool { _ = "STUB: not implemented"; return false }

func (q *ArrayQueue[T]) Print() { _ = "STUB: not implemented"; return }
