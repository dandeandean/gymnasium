package main

import "fmt"

// Given
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func simpleSwap(cur *TreeNode) {
	oldLeft := cur.Left
	cur.Left = cur.Right
	cur.Right = oldLeft
}
func invertTree(root *TreeNode) *TreeNode {
	if nil == root || (nil == root.Left && nil == root.Right) {
		return root
	}
	simpleSwap(root)
	if nil != root.Left {
		invertTree(root.Left)
	}
	if nil != root.Right {
		invertTree(root.Right)
	}
	return root
}
func main() {
	fmt.Print("HLW")
}
