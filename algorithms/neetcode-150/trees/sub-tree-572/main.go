package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeCmp(t0 *TreeNode, t1 *TreeNode) bool {
	if t0 == nil && t1 == nil {
		return true
	}
	if t0 == nil || t1 == nil || t0.Val != t1.Val {
		return false
	}
	return treeCmp(t0.Left, t1.Left) && treeCmp(t0.Right, t1.Right)

}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val {
		if treeCmp(root, subRoot) {
			return true
		}
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
func main() {
}
