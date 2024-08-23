package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getWeight(n *TreeNode) int {
	if nil == n {
		return 0
	}
	return 1 + max(getWeight(n.Right), getWeight(n.Left))
}
func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	leftWeight := getWeight(root.Left)
	rightWeight := getWeight(root.Right)
	if leftWeight > rightWeight {

		return leftWeight-rightWeight <= 1
	}
	return rightWeight-leftWeight <= 1
}

func main() {
}
