package datastructure

type LinkNode[T any] struct {
	Value T
	Pre   *LinkNode[T]
	Next  *LinkNode[T]
}

func NewLinkNode[T any](value T) *LinkNode[T] { _ = "STUB: not implemented"; return nil }

type StackNode[T any] struct {
	Value T
	Next  *StackNode[T]
}

func NewStackNode[T any](value T) *StackNode[T] { _ = "STUB: not implemented"; return nil }

type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}

func NewQueueNode[T any](value T) *QueueNode[T] { _ = "STUB: not implemented"; return nil }

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T any](val T) *TreeNode[T] { _ = "STUB: not implemented"; return nil }
