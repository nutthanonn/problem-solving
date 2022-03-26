// เเนวคิด

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

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
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

func mergeTwoLists(l1, l2 *linkedlist) {
	arr := []int{}
	ptr1 := l1.head
	ptr2 := l2.head
	for i := 0; i < l1.size; i++ {
		arr = append(arr, ptr1.Val)
		ptr1 = ptr1.next
	}
	for i := 0; i < l2.size; i++ {
		arr = append(arr, ptr2.Val)
		ptr2 = ptr2.next
	}
	arr = mergeSort(arr)
	l1.head = nil
	l1.size = 0
	for i := 0; i < len(arr); i++ {
		l1.Insert(arr[i])
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
	l1 := linkedlist{}
	l2 := linkedlist{}

	l1.Insert(1)
	l1.Insert(2)
	l1.Insert(4)

	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	mergeTwoLists(&l1, &l2)
	l1.DataList()

}
