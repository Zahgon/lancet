package datastructure

type CircularQueue[T any] struct {
	data     []T
	front    int
	rear     int
	capacity int
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] { _ = "STUB: not implemented"; return nil }

func (q *CircularQueue[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (q *CircularQueue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (q *CircularQueue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *CircularQueue[T]) IsFull() bool { _ = "STUB: not implemented"; return false }

func (q *CircularQueue[T]) Front() T { _ = "STUB: not implemented"; return *new(T) }

func (q *CircularQueue[T]) Back() T { _ = "STUB: not implemented"; return *new(T) }

func (q *CircularQueue[T]) Enqueue(value T) error { _ = "STUB: not implemented"; return nil }

func (q *CircularQueue[T]) Dequeue() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *CircularQueue[T]) Clear() { _ = "STUB: not implemented"; return }

func (q *CircularQueue[T]) Contain(value T) bool { _ = "STUB: not implemented"; return false }

func (q *CircularQueue[T]) Print() { _ = "STUB: not implemented"; return }
