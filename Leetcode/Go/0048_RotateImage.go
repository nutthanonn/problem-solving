package main

import "fmt"

func rotate(matrix [][]int) {
	arr := [][]int{}
	for i := 0; i < len(matrix); i++ {
		reverseArr := []int{}
		for j := len(matrix) - 1; j > -1; j-- {
			reverseArr = append(reverseArr, matrix[j][i])
		}
		arr = append(arr, reverseArr)
	}
	matrix = arr
}

func main() {
	number := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(number)
	fmt.Println(number)
}
