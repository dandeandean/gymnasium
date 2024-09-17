package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{root}
	cur := root
	visited := 0
	for len(stack) > 0 {
		//go left
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited++
		if visited == k {
			break
		}
		// go right
		cur = cur.Right
	}
	return cur.Val
}

func main() {

}
