package main

import "fmt"

func findMin(v []int) int {
	m := v[0]
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}

func findIndex(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func luckyNumbers(matrix [][]int) []int {
	arr := []int{}
	for i := 0; i < len(matrix); i++ {
		min := findMin(matrix[i])
		index := findIndex(matrix[i], min)
		newArr := []int{}
		for j := 0; j < len(matrix); j++ {
			newArr = append(newArr, matrix[j][index])
		}
		ans := findMax(newArr)
		if ans == min {
			arr = append(arr, ans)
		}
	}
	return arr
}

func main() {
	matrix := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	fmt.Println(luckyNumbers(matrix))
}
