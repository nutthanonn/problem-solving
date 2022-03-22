//   เเนวคิดเฉยๆ

package main

import (
	"math"
	"strconv"
)

type Node struct {
	next *Node
	Val  int
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
		if ptr.next == nil {
			ptr.next = &new_node
			l.size++
			return
		}
		ptr = ptr.next
	}
}

func addTwoNumbers(l1, l2 linkedlist) []string {
	ptr1 := l1.head
	ptr2 := l2.head
	seq1 := 0
	seq2 := 0

	for i := 0; i < l1.size; i++ {
		seq1 += ptr1.Val * int(math.Pow(10, float64(i)))
		ptr1 = ptr1.next
	}
	for i := 0; i < l2.size; i++ {
		seq2 += ptr2.Val * int(math.Pow(10, float64(i)))
		ptr2 = ptr2.next
	}

	temp := strconv.Itoa(seq1 + seq2)
	arr := []string{}
	for i := len(temp) - 1; i > -1; i-- {
		arr = append(arr, string(temp[i]))
	}

	return arr
}

func main() {
	l1 := linkedlist{}
	l2 := linkedlist{}

	l1.Insert(2)
	l1.Insert(4)
	l1.Insert(3)

	l2.Insert(5)
	l2.Insert(6)
	l2.Insert(4)

	addTwoNumbers(l1, l2)

}
