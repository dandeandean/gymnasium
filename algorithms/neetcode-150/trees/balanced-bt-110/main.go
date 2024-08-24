package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type TreeTuple struct {
	Height   int
	Balanced bool
}

func dfsBalanced(n *TreeNode) TreeTuple {
	if nil == n {
		return TreeTuple{0, true}
	}
	leftWeight := dfsBalanced(n.Left)
	rightWeight := dfsBalanced(n.Right)
	diff := abs(leftWeight.Height - rightWeight.Height)
	bal := leftWeight.Balanced && rightWeight.Balanced && (diff <= 1)
	return TreeTuple{
		Height:   1 + max(leftWeight.Height, rightWeight.Height),
		Balanced: bal,
	}
}
func isBalanced(root *TreeNode) bool {
	return dfsBalanced(root).Balanced
}

func main() {
}
