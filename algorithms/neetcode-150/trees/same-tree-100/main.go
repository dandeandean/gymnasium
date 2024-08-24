package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p != nil) && (q != nil) {
		if p.Val != q.Val {
			return false
		}
		left := isSameTree(p.Left, q.Left)
		right := isSameTree(p.Right, q.Right)
		if !left || !right {
			return false
		}
	}
	if ((p != nil) && (q == nil)) || ((q != nil) && (p == nil)) {
		return false
	}
	return true
}

func main() {
	fmt.Print("Not this again")
}
