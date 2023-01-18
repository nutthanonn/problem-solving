package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if arr[i]%2 != 0 && arr[i+1]%2 != 0 && arr[i+2]%2 != 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(threeConsecutiveOdds([]int{2, 6, 4, 1}))
	fmt.Println(threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
}
