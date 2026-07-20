package iterator

func Map[T any, U any](iter Iterator[T], iteratee func(item T) U) Iterator[U] {
	_ = "STUB: not implemented"
	return nil
}

type mapIterator[T any, U any] struct {
	iter     Iterator[T]
	iteratee func(T) U
}

func (mr *mapIterator[T, U]) HasNext() bool { _ = "STUB: not implemented"; return false }

func (mr *mapIterator[T, U]) Next() (U, bool) { _ = "STUB: not implemented"; return *new(U), false }

func Filter[T any](iter Iterator[T], predicateFunc func(item T) bool) Iterator[T] {
	_ = "STUB: not implemented"
	return nil
}

type filterIterator[T any] struct {
	iter          Iterator[T]
	predicateFunc func(T) bool
}

func (fr *filterIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (fr *filterIterator[T]) HasNext() bool { _ = "STUB: not implemented"; return false }

func Join[T any](iters ...Iterator[T]) Iterator[T] { _ = "STUB: not implemented"; return nil }

type joinIterator[T any] struct {
	iters []Iterator[T]
}

func (iter *joinIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (iter *joinIterator[T]) HasNext() bool { _ = "STUB: not implemented"; return false }

func Reduce[T any, U any](iter Iterator[T], initial U, reducer func(U, T) U) U {
	_ = "STUB: not implemented"
	return *new(U)
}

func Take[T any](it Iterator[T], num int) Iterator[T] { _ = "STUB: not implemented"; return nil }

type takeIterator[T any] struct {
	it  Iterator[T]
	num int
}

func (iter *takeIterator[T]) Next() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (iter *takeIterator[T]) HasNext() bool { _ = "STUB: not implemented"; return false }
