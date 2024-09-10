package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, goodest int) int {
	if node == nil {
		return 0
	}
	if node.Val >= goodest {
		return dfs(node.Left, node.Val) + dfs(node.Right, node.Val) + 1
	}
	return dfs(node.Left, goodest) + dfs(node.Right, goodest)
}
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	goodnessL := dfs(root.Left, root.Val)
	goodnessR := dfs(root.Right, root.Val)
	return goodnessL + goodnessR + 1
}
func main() {}
