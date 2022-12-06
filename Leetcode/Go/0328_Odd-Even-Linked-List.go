package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(link *ListNode, val int) {
	temp := link
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: val}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func printLinked(linked *ListNode) {
	temp := linked
	for temp.Next != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
	fmt.Println(temp.Val)
}

func main() {
	nums := []int{2, 1, 3, 5, 6, 4, 7}
	h := &ListNode{}
	for _, v := range nums {
		addNode(h, v)
	}

	res := oddEvenList(h.Next)

	printLinked(res)
}
