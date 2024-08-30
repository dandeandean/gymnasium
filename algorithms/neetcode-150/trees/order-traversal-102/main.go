package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var out [][]int
	q := make([]*TreeNode, 0)
	if root != nil {
		q = append(q, root)
	}
	for len(q) != 0 {
		vals := make([]int, 0)
		for i := 0; i < len(q); i++ {
			node := q[0] //pop.step.0
			q = q[1:]    //pop.step.1
			vals = append(vals, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			out = append(out, vals)
		}

	}
	return out

}
func main() {
	fmt.Println("Hallo Leute!")
}
