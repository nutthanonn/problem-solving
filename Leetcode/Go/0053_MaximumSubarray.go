package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {
	sum, currentSum := nums[0], nums[0]
	for _, v := range nums[1:] {
		sum = max(v, sum+v)
		currentSum = max(sum, currentSum)
	}
	return currentSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
