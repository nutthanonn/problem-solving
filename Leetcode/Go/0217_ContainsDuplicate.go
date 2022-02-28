package main

import "fmt"

func containsDuplicate(nums []int) bool {
	newArr := make(map[int]int)
	for _, v := range nums {
		newArr[v]++
	}
	if len(nums) != len(newArr) {
		return true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
