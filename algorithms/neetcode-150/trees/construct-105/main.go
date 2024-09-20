package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := new(TreeNode)
	if len(preorder) != len(inorder) {
		return root
	}
	// PreOrder := root + left sub tree + right sub tree
	// InOrder : = left sub + node + right sub
	// Because of this we can start with the PreOrder[0] as root
	root.Val = preorder[0]
	// Then use inorder's left sub tree
	cur := root
	for val := range inorder {
		if val == root.Val {
			break
		}
		newNode := new(TreeNode)
		newNode.Val = val
		cur.Left = newNode
		cur = cur.Left
	}
	return root
}
func main() {
}
