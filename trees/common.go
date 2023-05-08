package trees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeHeight returns the height of a binary  tree
func TreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lheight := TreeHeight(root.Left)
	rheight := TreeHeight(root.Right)

	if lheight > rheight {
		return lheight + 1
	}
	return rheight + 1
}

func PrintLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	height := TreeHeight(root)

	for i := 0; i < height; i++ {

	}
}
