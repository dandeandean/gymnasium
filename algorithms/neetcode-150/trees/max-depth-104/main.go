package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	leftDep := 1 + maxDepth(root.Left)
	rightDep := 1 + maxDepth(root.Right)
	if leftDep > rightDep {
		return leftDep
	}
	return rightDep
}
func main() {
	t := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	dep := maxDepth(t)
	fmt.Println("Depth = ", dep)
}
