package countcompletetreenodes

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    l := countNodes(root.Left)
    r := countNodes(root.Right)

    return 1 + l + r
}