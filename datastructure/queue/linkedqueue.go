package datastructure

import (
	"github.com/duke-git/lancet/v2/datastructure"
)

type LinkedQueue[T any] struct {
	head   *datastructure.QueueNode[T]
	tail   *datastructure.QueueNode[T]
	length int
}

func NewLinkedQueue[T any]() *LinkedQueue[T] { _ = "STUB: not implemented"; return nil }

func (q *LinkedQueue[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (q *LinkedQueue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (q *LinkedQueue[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *LinkedQueue[T]) Enqueue(value T) { _ = "STUB: not implemented"; return }

func (q *LinkedQueue[T]) Dequeue() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *LinkedQueue[T]) Front() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *LinkedQueue[T]) Back() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *LinkedQueue[T]) Clear() { _ = "STUB: not implemented"; return }

func (q *LinkedQueue[T]) Print() { _ = "STUB: not implemented"; return }

func (q *LinkedQueue[T]) Contain(value T) bool { _ = "STUB: not implemented"; return false }
