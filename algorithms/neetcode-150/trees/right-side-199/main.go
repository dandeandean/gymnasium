package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	out := make([]int, 0)
	for root != nil {
		out = append(out, root.Val)
		root = root.Right
	}
	return out
}
