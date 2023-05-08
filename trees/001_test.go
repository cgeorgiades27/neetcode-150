package trees

import (
	"fmt"
	"testing"
)

// Given the root of a binary tree, invert the tree, and return its root.

func InvertTree(root *TreeNode) *TreeNode {
	if root.Left == nil || root.Right == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	InvertTree(root.Left)
	InvertTree(root.Right)
	return nil
}

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func TestInvertTree(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println("before")
	preorderTraversal(root)
	InvertTree(root)
	fmt.Println("after")
	preorderTraversal(root)
}
