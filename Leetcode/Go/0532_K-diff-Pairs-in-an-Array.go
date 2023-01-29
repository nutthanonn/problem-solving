package main

import "fmt"

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || k < 0 {
		return 0
	}

	m := make(map[int]int)
	c := 0
	for _, v := range nums {
		m[v]++
	}

	for key, val := range m {
		if k == 0 {
			if val >= 2 {
				c++
			}
		} else {
			if _, ok := m[key+k]; ok {
				c++
			}
		}
	}

	return c
}

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
}
