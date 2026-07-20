package datastructure

import (
	"github.com/duke-git/lancet/v2/datastructure"
)

type DoublyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}

func NewDoublyLink[T any]() *DoublyLink[T] { _ = "STUB: not implemented"; return nil }

func (dl *DoublyLink[T]) InsertAtHead(value T) { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) InsertAtTail(value T) { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) InsertAt(index int, value T) { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) DeleteAtHead() { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) DeleteAtTail() { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) DeleteAt(index int) { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) Reverse() { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) GetMiddleNode() *datastructure.LinkNode[T] {
	_ = "STUB: not implemented"
	return nil
}

func (dl *DoublyLink[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (dl *DoublyLink[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (dl *DoublyLink[T]) Print() { _ = "STUB: not implemented"; return }

func (dl *DoublyLink[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (dl *DoublyLink[T]) Clear() { _ = "STUB: not implemented"; return }
