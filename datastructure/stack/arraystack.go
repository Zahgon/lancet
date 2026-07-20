package datastructure

type ArrayStack[T any] struct {
	data   []T
	length int
}

func NewArrayStack[T any]() *ArrayStack[T] { _ = "STUB: not implemented"; return nil }

func (s *ArrayStack[T]) Data() []T { _ = "STUB: not implemented"; return nil }

func (s *ArrayStack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *ArrayStack[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *ArrayStack[T]) Push(value T) { _ = "STUB: not implemented"; return }

func (s *ArrayStack[T]) Pop() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *ArrayStack[T]) Peak() (*T, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *ArrayStack[T]) Clear() { _ = "STUB: not implemented"; return }
