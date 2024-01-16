package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int

	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}
