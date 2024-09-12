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

func checker(father *TreeNode, lBound, rBound int) bool {
	if father == nil {
		return true
	}
	if (father.Val <= lBound) || (father.Val >= rBound) {
		return false
	}

	return checker(father.Left, lBound, father.Val) && checker(father.Right, father.Val, rBound)
}
func isValidBST(root *TreeNode) bool {
	leftBound := math.MinInt
	rightBound := math.MaxInt
	return checker(root, leftBound, rightBound)
}
func main() {
	fmt.Println("Validating...")
}
