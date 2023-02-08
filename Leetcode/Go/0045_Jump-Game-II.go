package main

import "fmt"

func jump(nums []int) int {
	minimum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		minimum[i] = len(nums)
	}

	minimum[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(nums) {
				if minimum[i+j] > minimum[i]+1 {
					minimum[i+j] = minimum[i] + 1
				}
			}
		}
	}

	return minimum[len(nums)-1]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
