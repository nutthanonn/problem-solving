package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func deleteNode(node *ListNode) {
	nextNode := node.Next
	node.Val = nextNode.Val
	node.Next = nextNode.Next
	nextNode.Next = nil
}

func (this *ListNode) printList() {
	for this != nil {
		fmt.Println(this.Val)
		this = this.Next
	}
}

func main() {
	l := &ListNode{}
	l.Val = 1
	l.Next.Val = 2

	l.printList()

}
