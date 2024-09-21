package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// PreOrder := root + left sub tree + right sub tree
	// InOrder : = left sub + node + right sub
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var leftPre, leftIn, rightPre, rightIn []int
	for i, val := range inorder {
		if val == preorder[0] {
			leftPre = preorder[1 : i+1]
			leftIn = inorder[:i+1]
			rightPre = preorder[i+1:]
			rightIn = inorder[i+1:]
			break
		}
	}
	root.Left = buildTree(leftPre, leftIn)
	root.Right = buildTree(rightPre, rightIn)
	return root
}
func main() {
}
