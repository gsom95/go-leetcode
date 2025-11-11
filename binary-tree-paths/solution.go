package binarytreepaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var path []string
	if root == nil {
		return path
	}

	curPath := strconv.FormatInt(int64(root.Val), 10)
	if root.Left == nil && root.Right == nil {
		return append(path, curPath)
	}

	leftPaths := binaryTreePaths(root.Left)
	rightPaths := binaryTreePaths(root.Right)

	for _, p := range leftPaths {
		path = append(path, curPath+"->"+p)
	}
	for _, p := range rightPaths {
		path = append(path, curPath+"->"+p)
	}

	return path
}
