package main

import "fmt"

func moveZeroes(nums []int) {
	final := []int{}
	c := 0
	for _, v := range nums {
		if v == 0 {
			c++
		} else {
			final = append(final, v)
		}
	}
	for i := 0; i < c; i++ {
		final = append(final, 0)
	}
	copy(nums, final)
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
