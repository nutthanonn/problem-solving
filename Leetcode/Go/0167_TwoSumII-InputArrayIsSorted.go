package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 1 {
		return numbers
	}
	first := 0
	last := len(numbers) - 1
	for first < last {
		val := numbers[first] + numbers[last]
		if val == target {
			return []int{first + 1, last + 1}
		} else if val < target {
			first++
		} else {
			last--
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
