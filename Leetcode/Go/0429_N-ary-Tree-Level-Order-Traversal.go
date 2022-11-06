package main

import "fmt"

type Node struct {
	children []*Node
	Val      int
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue []*Node
	queue = append(queue, root)

	for len(queue) > 0 {
		var level []int
		var next []*Node

		for _, node := range queue {
			level = append(level, node.Val)
			next = append(next, node.children...)
		}

		result = append(result, level)
		queue = next
	}
	return result
}

func main() {
	root := &Node{Val: 3}
	root.children = append(root.children, &Node{Val: 5})
	root.children = append(root.children, &Node{Val: 6})
	root.children[0].children = append(root.children[0].children, &Node{Val: 2})
	root.children[0].children = append(root.children[0].children, &Node{Val: 4})

	fmt.Println(levelOrder(root))
}
