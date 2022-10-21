package main

import "fmt"

func abs(nums int) int {
	if nums < 0 {
		return -nums
	}
	return nums
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok && abs(i-m[v]) <= k {
			return true
		}
		m[v] = i
	}
	return false
}

func main() {
	a := containsNearbyDuplicate([]int{1, 0, 1, 1, 2, 3}, 2)
	fmt.Println(a)
}
