package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	center := len(nums) / 2
	if len(nums) == 0 {
		return float64(0)
	} else if len(nums) == 1 {
		return float64(nums[0])
	}

	if len(nums)%2 != 0 {
		return float64(nums[center])
	} else {
		return (float64(nums[center-1]) + float64(nums[center])) / 2
	}

}

func main() {
	nums1 := []int{}
	nums2 := []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
