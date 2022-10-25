package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
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

func (root *TreeNode) maxLevelSum() int {

	q := []*TreeNode{}
	q = append(q, root)

	min_val := root.Val
	level_count := 0
	res := 1

	for len(q) > 0 {
		count := len(q)
		val := 0
		for count > 0 {
			temp := q[0]
			q = q[1:]
			val += temp.Val
			if temp.Left != nil {
				q = append(q, temp.Left)
			}

			if temp.Right != nil {
				q = append(q, temp.Right)
			}
			count--
		}
		if val > min_val {
			min_val = val
			res = level_count
		}
		level_count += 1
	}

	return res
}

func main() {
	t := &TreeNode{Val: 10}
	t.Insert(20)
	t.Insert(-10)
	t.Insert(3)
	t.Insert(100)
	fmt.Println(t.maxLevelSum())
}
