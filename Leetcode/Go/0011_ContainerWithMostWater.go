package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	mostWater := 0

	for i < j {
		mostWater = max(min(height[i], height[j])*(j-i), mostWater)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return mostWater
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
