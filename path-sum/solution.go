package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	hasLeft := hasPathSum(root.Left, targetSum-root.Val)
	hasRight := hasPathSum(root.Right, targetSum-root.Val)

	return hasLeft || hasRight
}
