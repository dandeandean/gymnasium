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
		level := make([]int, 0)
		lenQ := len(q)
		for i := 0; i < lenQ; i++ {
			youAreHere := q[0]
			q = q[1:]
			if youAreHere == nil {
				break
			}
			level = append(level, youAreHere.Val)
			if youAreHere.Left != nil {
				q = append(q, youAreHere.Left)
			}
			if youAreHere.Right != nil {
				q = append(q, youAreHere.Right)
			}
		}
		if len(level) != 0 {
			out = append(out, level[len(level)-1])
		}
	}
	return out
}
