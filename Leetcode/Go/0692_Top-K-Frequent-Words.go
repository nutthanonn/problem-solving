package main

import "fmt"

type Listnode struct {
	next   *Listnode
	val    string
	weight int
}

func (this *Listnode) Insert(val string) {
	curr := this
	for curr.next != nil {
		if curr.val == val {
			curr.weight += 1
			return
		}
		curr = curr.next
	}

	if curr.val == val {
		curr.weight += 1
		return
	}
	curr.next = &Listnode{val: val, weight: 1}
}

func (this *Listnode) InsertSort(val string, weight, k int) {
	curr := this
	for i := 0; i < k; i++ {
		if curr.next.weight < weight {
			temp := &Listnode{val: val, weight: weight}
			temp.next = curr.next
			curr.next = temp
			return
		}

		if curr.next.weight == weight {
			if val < curr.next.val {
				temp := &Listnode{val: val, weight: weight}
				temp.next = curr.next
				curr.next = temp
				return
			}
		}

		curr = curr.next
	}
}

func topKFrequent(words []string, k int) []string {
	h := &Listnode{}
	for _, v := range words {
		h.Insert(v)
	}
	// insertSort
	sorted_kNode := &Listnode{}

	curr := sorted_kNode
	for i := 0; i < k; i++ {
		if curr.next == nil {
			curr.next = &Listnode{val: "", weight: -1}
		}
		curr = curr.next
	}

	temp := h
	for temp != nil {
		sorted_kNode.InsertSort(temp.val, temp.weight, k)
		temp = temp.next
	}
	sorted_kNode = sorted_kNode.next
	out := []string{}
	for i := 0; i < k; i++ {
		out = append(out, sorted_kNode.val)
		sorted_kNode = sorted_kNode.next
	}

	return out
}

func (this *Listnode) print() {
	curr := this.next
	for curr != nil {
		fmt.Println(curr.val, curr.weight)
		curr = curr.next
	}
}

func main() {
	a := topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 3)
	fmt.Println(a)

}
