package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//  v - LeetcodeSolution - v
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	count := 1

	if ptr.Next == nil && n == 1 {
		return nil
	}

	for ptr.Next != nil {
		ptr = ptr.Next
		count++
	}
	count = count - n

	ptr = head
	if count == 0 {
		head = head.Next
		return head
	}

	for i := 0; i < count-1; i++ {
		ptr = ptr.Next
	}

	prevNode := ptr
	prevNode.Next = ptr.Next.Next

	return head
}

//  ^ - LeetcodeSolution - ^

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
		fmt.Println(ptr.Val)
	}
}

func main() {
	l := &ListNode{Val: 1}
	l.Insert(2)
	// l.Insert(3)
	// l.Insert(4)
	// l.Insert(5)

	a := removeNthFromEnd(l, 2)
	a.Search()
}
