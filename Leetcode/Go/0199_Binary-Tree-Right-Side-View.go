package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (this *TreeNode) Insert(val int) {
	new_node := &TreeNode{}
	new_node.Val = val

	if val > this.Val {
		if this.Right == nil {
			this.Right = new_node
		} else {
			this.Right.Insert(val)
		}
	} else if val < this.Val {
		if this.Left == nil {
			this.Left = new_node
		} else {
			this.Left.Insert(val)
		}
	}
}

func (this *TreeNode) printLevelOrderLinebyLine() []int {
	if this == nil {
		return nil
	}

	q := []*TreeNode{}
	q = append(q, this)
	col := []int{this.Val}
	for len(q) != 0 {
		count := len(q)
		for count > 0 {
			temp := q[0]
			q = q[1:]
			// fmt.Printf("%d  ", temp.Val)
			if temp.Left != nil {
				q = append(q, temp.Left)
			}

			if temp.Right != nil {
				q = append(q, temp.Right)
			}
			count--
		}
		if len(q) != 0 {
			col = append(col, q[len(q)-1].Val)
		}
	}
	return col
}

func main() {
	t := &TreeNode{Val: 100}
	t.Insert(150)
	t.Insert(50)
	t.Insert(25)

	fmt.Println(t.printLevelOrderLinebyLine())
}
