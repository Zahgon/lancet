package promise

import (
	"sync"
)

type Promise[T any] struct {
	runnable func(resolve func(T), reject func(error))
	result   T
	err      error

	pending bool

	mu *sync.Mutex
	wg *sync.WaitGroup
}

func New[T any](runnable func(resolve func(T), reject func(error))) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Promise[T]) run() { _ = "STUB: not implemented"; return }

func Resolve[T any](resolution T) *Promise[T] { _ = "STUB: not implemented"; return nil }

func (p *Promise[T]) resolve(value T) { _ = "STUB: not implemented"; return }

func Reject[T any](err error) *Promise[T] { _ = "STUB: not implemented"; return nil }

func (p *Promise[T]) reject(err error) { _ = "STUB: not implemented"; return }

func Then[T1, T2 any](promise *Promise[T1], resolve1 func(value T1) T2) *Promise[T2] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Promise[T]) Then(resolve func(value T) T) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

func Catch[T any](promise *Promise[T], rejection func(err error) error) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Promise[T]) Catch(reject func(error) error) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Promise[T]) Await() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

type tuple[T1, T2 any] struct {
	_1 T1
	_2 T2
}

func All[T any](promises []*Promise[T]) *Promise[[]T] { _ = "STUB: not implemented"; return nil }

func Race[T any](promises []*Promise[T]) *Promise[T] { _ = "STUB: not implemented"; return nil }

func Any[T any](promises []*Promise[T]) *Promise[T] { _ = "STUB: not implemented"; return nil }
