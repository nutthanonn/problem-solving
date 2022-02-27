package main

import "fmt"

func IsIn(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}

func searchInsert(nums []int, target int) int {
	if IsIn(nums, target) {
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] > target {
				return i
			}
		}
	}
	return len(nums)
}

// func searchInsert(nums []int, target int) int {
// 	var j, i int
// 	for j, i = range nums {
// 		if i >= target {
// 			return j
// 		}
// 	}
// 	j++
// 	return j
// }

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))

}
