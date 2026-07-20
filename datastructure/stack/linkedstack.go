package datastructure

import (
	"github.com/duke-git/lancet/v2/datastructure"
)

type LinkedStack[T any] struct {
	top    *datastructure.StackNode[T]
	length int
}

func NewLinkedStack[T any]() *LinkedStack[T] { _ = "STUB: not implemented"; return nil }

func (s *LinkedStack[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (s *LinkedStack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *LinkedStack[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *LinkedStack[T]) Push(value T) { _ = "STUB: not implemented"; return }

func (s *LinkedStack[T]) Pop() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *LinkedStack[T]) Peak() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *LinkedStack[T]) Clear() { _ = "STUB: not implemented"; return }

func (s *LinkedStack[T]) Print() { _ = "STUB: not implemented"; return }
