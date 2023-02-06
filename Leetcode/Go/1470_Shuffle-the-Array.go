package main

import "fmt"

func shuffle(nums []int, n int) []int {
	final := []int{}

	for i := 0; i < len(nums)/2; i += 1 {
		final = append(final, nums[i])
		final = append(final, nums[(len(nums)/2)+i])
	}

	return final
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}
