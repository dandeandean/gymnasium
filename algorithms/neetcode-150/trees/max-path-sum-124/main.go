package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(node *TreeNode, sum *int) int {
	if node == nil {
		return 0
	}
	leftSide := max(dfs(node.Left, sum), 0)
	rightSide := max(dfs(node.Right, sum), 0)
	*sum = max(*sum, node.Val+leftSide+rightSide)
	return node.Val + max(leftSide, rightSide)
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val
	dfs(root, &sum)
	return sum
}

func main() {
}
