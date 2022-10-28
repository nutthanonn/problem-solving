package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) insert(val int) {
	if this.Val > val {
		if this.Left == nil {
			this.Left = &TreeNode{Val: val}
		} else {
			this.Left.insert(val)
		}
	} else {
		if this.Right == nil {
			this.Right = &TreeNode{Val: val}
		} else {
			this.Right.insert(val)
		}
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func (this *TreeNode) levelOrder() {
	if this == nil {
		return
	}
	queue := []*TreeNode{this}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	t := &TreeNode{Val: 100}
	t.insert(50)
	t.insert(150)
	t.insert(25)
	t.insert(75)
	t.insert(125)
	t.insert(175)

	invertTree(t)

}
