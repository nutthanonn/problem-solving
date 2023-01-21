package main

import (
	"fmt"
	"math"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func constructRectangle(area int) []int {
	if area == 1 {
		return []int{1, 1}
	}

	min_diff := math.MaxInt
	final := []int{0, 0}
	for i := 1; i < int(math.Sqrt(float64(area)))+1; i++ {
		if area%i == 0 {
			v := []int{i, area / i}
			temp := abs(v[0] - v[1])
			if temp < min_diff {
				if v[1] < v[0] {
					final = []int{v[0], v[1]}
				} else {
					final = []int{v[1], v[0]}
				}
				min_diff = temp
			}
		}
	}

	return final
}

func main() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(37))
	fmt.Println(constructRectangle(122122))
}
