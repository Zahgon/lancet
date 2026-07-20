package datastructure

import (
	"github.com/duke-git/lancet/v2/datastructure"
)

type SinglyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}

func NewSinglyLink[T any]() *SinglyLink[T] { _ = "STUB: not implemented"; return nil }

func (sl *SinglyLink[T]) InsertAtHead(value T) { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) InsertAtTail(value T) { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) InsertAt(index int, value T) { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) DeleteAtHead() { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) DeleteAtTail() { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) DeleteAt(index int) { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) DeleteValue(value T) { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) Reverse() { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) GetMiddleNode() *datastructure.LinkNode[T] {
	_ = "STUB: not implemented"
	return nil
}

func (sl *SinglyLink[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (sl *SinglyLink[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (sl *SinglyLink[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (sl *SinglyLink[T]) Clear() { _ = "STUB: not implemented"; return }

func (sl *SinglyLink[T]) Print() { _ = "STUB: not implemented"; return }
