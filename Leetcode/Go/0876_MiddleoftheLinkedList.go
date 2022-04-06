package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//  v - LeetcodeSolution - v
func middleNode(head *ListNode) *ListNode {

	//get size of linkedlist
	temp := head
	c := 0
	for temp != nil {
		temp = temp.Next
		c++
	}

	for i := 0; i < c/2; i++ {
		head = head.Next
	}

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
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Search()

	fmt.Println(middleNode(l))

}
