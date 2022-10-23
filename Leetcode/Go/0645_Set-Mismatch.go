package main

import "fmt"

func solve(nums []int) []int {
	final := make(map[int]int)
	for _, v := range nums {
		if _, ok := final[v]; !ok {
			final[v] = 1
		} else {
			final[v]++
		}
	}
	result := []int{}
	for k, v := range final {
		if v == 2 {
			result = append(result, k)
		}
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, ok := final[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	fmt.Println(solve([]int{1, 1, 3, 2}))

}
