package main

import "fmt"

func runningSum(nums []int) []int {
	final := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		final = append(final, final[i-1]+nums[i])
	}

	return final
}

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}
