package main

import "fmt"

func rotate(nums []int, k int) {
	final := []int{}

	for i := 0; i < k; i++ {
		final = append(nums[len(nums)-1:], nums...)
		final = final[:len(final)-1]
	}
	copy(nums, final)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
