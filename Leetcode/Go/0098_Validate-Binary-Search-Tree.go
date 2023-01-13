package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(inorder(root.Left), root.Val), inorder(root.Right)...)
}

func isValidBST(root *TreeNode) bool {
	final := inorder(root)
	for i := 0; i < len(final)-1; i++ {
		if final[i] >= final[i+1] {
			return false
		}
	}

	return true
}

func main() {

}
