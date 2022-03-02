package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {

	index := 0
	len1 := len(nums1)
	len2 := len(nums2)

	for i := 0; i < len1; i++ {
		for j := index; j < len2; j++ {

			if nums1[i] == nums2[j] {
				nums1[index] = nums1[i]
				nums2[index], nums2[j] = nums2[j], nums2[index]
				index++
				break
			}

		}
	}

	return nums1[:index]
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}
