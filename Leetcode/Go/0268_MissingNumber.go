package main

import "fmt"

func missingNumber(nums []int) int {
	c := 0
	for _, v := range nums {
		c += v
	}
	all := len(nums) * (len(nums) + 1) / 2
	return all - c
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
