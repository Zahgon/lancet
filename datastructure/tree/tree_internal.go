package datastructure

import (
	"github.com/duke-git/lancet/v2/constraints"
	"github.com/duke-git/lancet/v2/datastructure"
)

func preOrderTraverse[T any](node *datastructure.TreeNode[T]) []T {
	_ = "STUB: not implemented"
	return nil
}

func postOrderTraverse[T any](node *datastructure.TreeNode[T]) []T {
	_ = "STUB: not implemented"
	return nil
}

func inOrderTraverse[T any](node *datastructure.TreeNode[T]) []T {
	_ = "STUB: not implemented"
	return nil
}

func levelOrderTraverse[T any](root *datastructure.TreeNode[T], traversal *[]T) {
	_ = "STUB: not implemented"
	return
}

func insertTreeNode[T any](rootNode, newNode *datastructure.TreeNode[T], comparator constraints.Comparator) {
	_ = "STUB: not implemented"
	return
}

func deleteTreeNode[T any](node *datastructure.TreeNode[T], data T, comparator constraints.Comparator) *datastructure.TreeNode[T] {
	_ = "STUB: not implemented"
	return nil
}

func inOrderSuccessor[T any](root *datastructure.TreeNode[T]) *datastructure.TreeNode[T] {
	_ = "STUB: not implemented"
	return nil
}

func printTreeNodes[T any](nodes []*datastructure.TreeNode[T], level, maxLevel int) {
	_ = "STUB: not implemented"
	return
}

func printSpaces(n int) { _ = "STUB: not implemented"; return }

func isAllNil[T any](nodes []*datastructure.TreeNode[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func calculateDepth[T any](node *datastructure.TreeNode[T], depth int) int {
	_ = "STUB: not implemented"
	return 0
}

func isSubTree[T any](superTreeRoot, subTreeRoot *datastructure.TreeNode[T], comparator constraints.Comparator) bool {
	_ = "STUB: not implemented"
	return false
}

func max(a, b int) int { _ = "STUB: not implemented"; return 0 }
