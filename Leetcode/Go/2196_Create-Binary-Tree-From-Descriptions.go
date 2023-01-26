package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func complement(slice [][]int) int {
	m := make(map[int]bool)
	for _, v := range slice {
		m[v[1]] = true
	}
	for _, v := range slice {
		if _, ok := m[v[0]]; !ok {
			return v[0]
		}
	}
	return -1
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	field := make(map[int]*TreeNode)

	for _, des := range descriptions {
		n := &TreeNode{Val: des[1]}

		if _, ok := field[des[1]]; !ok {
			field[des[1]] = &TreeNode{Val: des[1]}
		}

		if _, ok := field[des[0]]; !ok {
			field[des[0]] = &TreeNode{Val: des[0]}
		}

		temp := field[des[0]]

		if _, ok1 := field[des[1]]; !ok1 {
			if l := des[2]; l == 1 {
				temp.Left = n
			} else {
				temp.Right = n
			}
		} else {
			if l := des[2]; l == 1 {
				temp.Left = field[des[1]]
			} else {
				temp.Right = field[des[1]]
			}
		}
	}

	return field[complement(descriptions)]
}

func main() {
	fmt.Println(createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}).Left.Left.Val)
}
