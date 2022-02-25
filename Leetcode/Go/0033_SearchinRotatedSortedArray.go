package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	for k, v := range nums {
		if target == v {
			return k
		}
	}
	return -1
}

func main() {
	nums := []int{1}
	target := 0
	fmt.Println(search(nums, target))

}
