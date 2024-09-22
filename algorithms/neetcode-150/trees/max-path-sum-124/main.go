package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, sum *int) int {
	if node == nil {
		return 0
	}
	leftSide := dfs(node.Left, sum)
	rightSide := dfs(node.Right, sum)
	if leftSide < 0 {
		leftSide = 0
	}
	if rightSide < 0 {
		rightSide = 0
	}
	res := node.Val + leftSide + rightSide
	if res < 0 {
		res = 0
	}

	extra := 0
	if leftSide > rightSide {
		extra = leftSide
	}
	extra = rightSide

	return node.Val + extra
}
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val
	dfs(root.Left, &sum)
	dfs(root.Right, &sum)
	return sum
}
func main() {
}
