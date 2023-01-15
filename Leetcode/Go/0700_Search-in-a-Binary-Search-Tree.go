package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}

	if root.Left != nil {
		l := searchBST(root.Left, val)
		if l != nil {
			return l
		}
	}

	if root.Right != nil {
		r := searchBST(root.Right, val)
		if r != nil {
			return r
		}

	}

	return nil
}

func main() {

}
