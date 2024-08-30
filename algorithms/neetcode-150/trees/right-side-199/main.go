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
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		out = append(out, root.Val)
		root = root.Right
	}
	return out
}
