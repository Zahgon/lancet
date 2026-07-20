package datastructure

import (
	"github.com/duke-git/lancet/v2/constraints"
	"github.com/duke-git/lancet/v2/datastructure"
)

type BSTree[T any] struct {
	root       *datastructure.TreeNode[T]
	comparator constraints.Comparator
}

func NewBSTree[T any](rootData T, comparator constraints.Comparator) *BSTree[T] {
	_ = "STUB: not implemented"
	return nil
}

func (t *BSTree[T]) Insert(data T) { _ = "STUB: not implemented"; return }

func (t *BSTree[T]) Delete(data T) { _ = "STUB: not implemented"; return }

func (t *BSTree[T]) NodeLevel(node *datastructure.TreeNode[T]) int {
	_ = "STUB: not implemented"
	return 0
}

func (t *BSTree[T]) PreOrderTraverse() []T { _ = "STUB: not implemented"; return nil }

func (t *BSTree[T]) PostOrderTraverse() []T { _ = "STUB: not implemented"; return nil }

func (t *BSTree[T]) InOrderTraverse() []T { _ = "STUB: not implemented"; return nil }

func (t *BSTree[T]) LevelOrderTraverse() []T { _ = "STUB: not implemented"; return nil }

func (t *BSTree[T]) Depth() int { _ = "STUB: not implemented"; return 0 }

func (t *BSTree[T]) HasSubTree(subTree *BSTree[T]) bool { _ = "STUB: not implemented"; return false }

func hasSubTree[T any](superTreeRoot, subTreeRoot *datastructure.TreeNode[T],
	comparator constraints.Comparator) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *BSTree[T]) Print() { _ = "STUB: not implemented"; return }
