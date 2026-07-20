package optional

import (
	"sync"
)

type Optional[T any] struct {
	value *T
	mu    *sync.RWMutex
}

func Default[T any]() Optional[T] { _ = "STUB: not implemented"; return nil }

func Of[T any](value T) Optional[T] { _ = "STUB: not implemented"; return nil }

func FromNillable[T any](value *T) Optional[T] { _ = "STUB: not implemented"; return nil }

func (o Optional[T]) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (o Optional[T]) IsNil() bool { _ = "STUB: not implemented"; return false }

func (o Optional[T]) IfNotNil(action func(value T)) { _ = "STUB: not implemented"; return }

func (o Optional[T]) IfNotNilOrElse(action func(value T), fallbackAction func()) {
	_ = "STUB: not implemented"
	return
}

func (o Optional[T]) Unwarp() T { _ = "STUB: not implemented"; return *new(T) }

func (o Optional[T]) OrElse(other T) T { _ = "STUB: not implemented"; return *new(T) }

func (o Optional[T]) OrElseGet(action func() T) T { _ = "STUB: not implemented"; return *new(T) }

func (o Optional[T]) OrElseTrigger(errorHandler func() error) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
