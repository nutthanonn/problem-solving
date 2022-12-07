package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insert(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insert(root.Right, val)
		}
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	collect := []int{}

	q := []*TreeNode{}
	q = append(q, root)
	collect = append(collect, root.Val)
	count := 0

	if root.Val < high && root.Val > low {
		count += root.Val
	}

	for len(q) != 0 {
		temp := q[0]
		q = q[1:]

		if temp.Left != nil {
			q = append(q, temp.Left)
			if temp.Left.Val < high && temp.Left.Val > low {
				count += temp.Left.Val
			}
		}

		if temp.Right != nil {
			q = append(q, temp.Right)

			if temp.Right.Val < high && temp.Right.Val > low {
				count += temp.Right.Val
			}
		}
	}

	return count + low + high
}

func main() {
	t := &TreeNode{Val: 10}
	insert(t, 5)
	insert(t, 15)

	insert(t, 3)
	insert(t, 7)

	insert(t, 13)
	insert(t, 18)

	insert(t, 1)
	insert(t, 6)

	fmt.Println(rangeSumBST(t, 6, 10))
}
