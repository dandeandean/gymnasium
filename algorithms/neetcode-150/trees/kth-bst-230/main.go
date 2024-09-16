package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	nums := make([]int, 0)
	stack := []*TreeNode{root}
	cur := root
	for len(stack) > 0 {
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		top := stack[0]
		stack = stack[1:]
		nums = append(nums, top.Val)
	}
	return nums[k-1]
}

func main() {

}
