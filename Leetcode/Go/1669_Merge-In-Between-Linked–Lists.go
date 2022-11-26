package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addList(l *ListNode, val int) {
	temp := l
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: val}
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	left := list1

	for i := 0; i < a-1; i++ {
		left = left.Next
	}

	right := left
	for i := a; i < b+1; i++ {
		right = right.Next
	}

	left.Next = list2
	temp := list2

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = right.Next
	return list1
}

func printLink(l *ListNode) {
	temp := l.Next
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}

func main() {
	list1 := &ListNode{}
	addList(list1, 1)
	addList(list1, 2)
	addList(list1, 3)
	addList(list1, 4)
	addList(list1, 5)
	addList(list1, 6)
	addList(list1, 42)

	list2 := &ListNode{Val: 999}
	addList(list2, 100)
	addList(list2, 200)
	addList(list2, 300)

	a := mergeInBetween(list1, 2, 3, list2)
	printLink(a)
}
