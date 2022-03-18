package main

import "fmt"

func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	number := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Println(search(number, target))
}
