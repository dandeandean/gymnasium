package main

import "fmt"

type TreeNode struct {
	/**
	 * Definition for a binary tree node.
	 */
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, dep *int) int {
	if nil == node {
		return 0
	}
	leftSize := dfs(node.Left, dep)
	rightSize := dfs(node.Right, dep)
	*dep = max(leftSize+rightSize, *dep)
	return max(rightSize, leftSize) + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
	// base case: there are no children
	var dep = 0
	dfs(root, &dep)
	return dep
}
func main() {
	t := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	dia := diameterOfBinaryTree(t)
	fmt.Printf("diameter = %d\n", dia)

}
