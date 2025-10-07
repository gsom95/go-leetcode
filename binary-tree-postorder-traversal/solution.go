package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := postorderTraversal(root.Left)
	res = append(res, postorderTraversal(root.Right)...)

	return append(res, root.Val)
}
