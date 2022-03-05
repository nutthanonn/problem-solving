package main

import "fmt"

func generate(numRows int) [][]int {
	arr := [][]int{{1}, {1, 1}}
	if numRows == 1 {
		return [][]int{{1}}
	} else {
		for i := 2; i < numRows; i++ {
			arr_row := []int{1}
			for j := 0; j < i-1; j++ {
				arr_row = append(arr_row, arr[i-1][j]+arr[i-1][j+1])
			}
			arr_row = append(arr_row, 1)
			arr = append(arr, arr_row)
		}
		return arr
	}
}

func main() {
	numRows := 30
	fmt.Println(generate((numRows)))
}
