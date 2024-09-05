package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, goodness int) int {
	if node == nil {
		return 0
	}
	if node.Val <= goodness {
		return 1 + dfs(node.Left, goodness) + dfs(node.Right, goodness)
	}
	return dfs(node.Left, node.Val) + dfs(node.Right, node.Val)
}
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	goodness := 1 // 1 bc root must be good
	goodnessL := dfs(root.Left, goodness)
	goodnessR := dfs(root.Right, goodness)
	return goodnessL + goodnessR
}
func main() {}
