package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if abs(lh-rh) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

	return false
}

func main() {

}
