package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checker(father *TreeNode) bool {
	if (father == nil) || (father.Left == nil && father.Right == nil) {
		return true
	}
	if father.Right == nil {
		return father.Val > father.Left.Val
	}
	if father.Left == nil {
		return father.Val < father.Right.Val
	}
	//(father.Val < father.Right.Val) && (father.Val > father.Left.Val)
	return checker(father.Left) && checker(father.Right)
}
func isValidBST(root *TreeNode) bool {
	leftBound := math.MinInt
	rightBound := math.MaxInt
	return checker(root.Left) && checker(root.Right)
}
func main() {
	fmt.Println("Validating...")
}
