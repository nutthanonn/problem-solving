package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	current := prev
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		if sum >= 10 {
			sum = 1
			continue
		}
		sum = 0
	}

	if sum == 1 {
		current.Next = &ListNode{Val: sum}
	}

	return prev.Next
}

func (l *ListNode) Insert(val int) {
	new_node := ListNode{Val: val}
	ptr := l
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = &new_node
}

func (l *ListNode) Search() {
	ptr := l
	fmt.Println(ptr.Val)
	for ptr.Next != nil {
		ptr = ptr.Next
	}
}
func main() {
	l1 := &ListNode{Val: 2}
	l1.Insert(4)
	l1.Insert(3)

	l2 := &ListNode{Val: 5}
	l2.Insert(6)
	l2.Insert(4)

	a := addTwoNumbers(l1, l2)
	a.Search()
}
