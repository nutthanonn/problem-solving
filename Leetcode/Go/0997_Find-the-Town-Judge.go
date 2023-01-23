package main

import "fmt"

func max(arr []int) (int, int) {
	temp := arr[0]
	target := 0
	for i, v := range arr {
		if v > temp {
			temp = v
			target = i
		}
	}
	return temp, target
}

func findJudge(n int, trust [][]int) int {
	arr := make([]int, len(trust)+1)

	for _, v := range trust {
		arr[v[1]-1]++
		arr[v[0]-1]--
	}

	max_val, target := max(arr)

	if max_val == n-1 && arr[target] == n-1 {
		return target + 1
	}

	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	fmt.Println(findJudge(4, [][]int{{1, 4}, {2, 4}, {3, 4}}))
	fmt.Println(findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}))
}
