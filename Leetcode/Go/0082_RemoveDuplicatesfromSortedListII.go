package main

import "fmt"

type Node struct {
	Val  int
	next *Node
}

type linkedlist struct {
	head *Node
	size int
}

func (l *linkedlist) Insert(val int) {
	new_node := Node{}
	new_node.Val = val

	if l.size == 0 {
		l.head = &new_node
		l.size++
		return
	}

	ptr := l.head
	for i := 0; i < l.size; i++ {
		if ptr.Val == val {
			return
		} else if ptr.next == nil {
			ptr.next = &new_node
			l.size++
			return
		}
		ptr = ptr.next
	}

}

func (l *linkedlist) DataList() {
	ptr := l.head
	for i := 0; i < l.size; i++ {
		fmt.Println(ptr.Val)
		ptr = ptr.next
	}
}

func main() {
	l := linkedlist{}
	l.Insert(123)
	l.Insert(1)
	l.Insert(3)
	l.Insert(425)
	l.Insert(45)
	l.Insert(3)
	l.Insert(3)
	l.Insert(3)

	l.DataList()
}
