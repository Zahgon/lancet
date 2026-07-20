package iterator

import (
	"context"

	"golang.org/x/exp/constraints"
)

type Iterator[T any] interface {
	HasNext() bool

	Next() (item T, ok bool)
}

type ResettableIterator[T any] interface {
	Iterator[T]

	Reset()
}

type StopIterator[T any] interface {
	Iterator[T]

	Stop()
}

type DeleteIterator[T any] interface {
	Iterator[T]

	Delete()
}

type SetIterator[T any] interface {
	Iterator[T]

	Set(v T)
}

type PrevIterator[T any] interface {
	Iterator[T]

	Prev()
}

func FromSlice[T any](slice []T) *SliceIterator[T] { _ = "STUB: not implemented"; return nil }

func ToSlice[T any](iter Iterator[T]) []T { _ = "STUB: not implemented"; return nil }

type SliceIterator[T any] struct {
	slice []T
	index int
}

func (iter *SliceIterator[T]) HasNext() bool { _ = "STUB: not implemented"; return false }

func (iter *SliceIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (iter *SliceIterator[T]) Prev() { _ = "STUB: not implemented"; return }

func (iter *SliceIterator[T]) Set(value T) { _ = "STUB: not implemented"; return }

func (iter *SliceIterator[T]) Reset() { _ = "STUB: not implemented"; return }

func FromRange[T constraints.Integer | constraints.Float](start, end, step T) *RangeIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

type RangeIterator[T constraints.Integer | constraints.Float] struct {
	start, end, step, current T
}

func (iter *RangeIterator[T]) HasNext() bool { _ = "STUB: not implemented"; return false }

func (iter *RangeIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (iter *RangeIterator[T]) Reset() { _ = "STUB: not implemented"; return }

func FromChannel[T any](channel <-chan T) *ChannelIterator[T] {
	_ = "STUB: not implemented"
	return nil
}

type ChannelIterator[T any] struct {
	channel <-chan T
}

func (iter *ChannelIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (iter *ChannelIterator[T]) HasNext() bool { _ = "STUB: not implemented"; return false }

func ToChannel[T any](ctx context.Context, iter Iterator[T], buffer int) <-chan T {
	_ = "STUB: not implemented"
	return nil
}
