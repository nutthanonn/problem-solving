package main

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])

	for _, v := range matrix {
		if v[0] <= target && v[n-1] >= target {
			if binarySearch(v, target) {
				return true
			}
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5
	println(searchMatrix(matrix, target))
}
