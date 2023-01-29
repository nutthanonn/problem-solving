package main

import (
	"fmt"
)

func dis(i, j int) int {
	if i < j {
		return j - i
	}
	return i - j
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) <= 1 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+indexDiff && j < len(nums); j++ {
			if dis(nums[i], nums[j]) <= valueDiff {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
