package main

import (
	"fmt"
	"math"
)

func prefixSum(item []int) []int {
	prefix := []int{}
	prefix = append(prefix, item[0])

	for i := 1; i < len(item); i++ {
		prefix = append(prefix, prefix[i-1]+item[i])
	}

	return prefix
}

func diff(nums, d int) int {
	if d == 0 {
		return nums
	}
	return nums / d
}

func minimumAverageDifference(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	minimum := math.MaxInt
	final := 0
	nums = prefixSum(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		left := diff(nums[i], (i + 1))
		right := diff(nums[n-1]-nums[i], n-i-1)

		ans := left - right

		if ans < 0 {
			ans *= -1
		}

		if ans < minimum {
			minimum = ans
			final = i
		}
	}

	return final
}

func main() {
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
}
