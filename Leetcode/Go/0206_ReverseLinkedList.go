package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		current, prev, current.Next = current.Next, current, prev
	}
	head = prev
	return head
}

func main() {

}
