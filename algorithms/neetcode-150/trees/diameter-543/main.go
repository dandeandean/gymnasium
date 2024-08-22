package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	/**
	 * Definition for a binary tree node.
	 */
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, dep *int, wg *sync.WaitGroup) int {
	defer wg.Done()
	if nil == node {
		return 0
	}
	wg.Add(2)
	leftSize := dfs(node.Left, dep, wg)
	rightSize := dfs(node.Right, dep, wg)
	*dep = max(leftSize+rightSize, *dep)
	return max(rightSize, leftSize) + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
	// base case: there are no children
	var dep = 0
	var wg sync.WaitGroup
	wg.Add(1)
	go dfs(root, &dep, &wg)
	wg.Wait()
	return dep
}
func main() {
	t := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	dia := diameterOfBinaryTree(t)
	fmt.Printf("diameter = %d\n", dia)

}
