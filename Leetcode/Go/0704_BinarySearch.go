package main

import "fmt"

func search(nums []int, target int) int {
	first := 0
	last := len(nums) - 1

	for first <= last {

		midIndex := int((first + last) / 2)
		midVal := nums[midIndex]
		if midVal == target {
			return midIndex
		} else if midVal < target {
			first = midIndex + 1
		} else {
			last = midIndex - 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	taget := 9
	fmt.Println(search(nums, taget))
}
