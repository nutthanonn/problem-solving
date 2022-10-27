package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) Insert(val int) {
	if this.Val > val {
		if this.Left == nil {
			this.Left = &TreeNode{Val: val}
		} else {
			this.Left.Insert(val)
		}
	} else {
		if this.Right == nil {
			this.Right = &TreeNode{Val: val}
		} else {
			this.Right.Insert(val)
		}
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}

	if (q == nil || p == nil) || q.Val != p.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	t1 := &TreeNode{Val: 1}
	t1.Insert(2)
	t1.Insert(3)
	t1.Insert(4)

	t2 := &TreeNode{Val: 1}
	t2.Insert(2)
	t2.Insert(3)
	t2.Insert(4)
}
