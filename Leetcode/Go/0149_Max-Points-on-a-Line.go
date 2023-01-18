package main

import "fmt"

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		if points[0][0] == points[0][1] {
			return 1
		}
	}

	max_path := 0
	for i := 0; i < len(points); i++ {
		start := points[i]
		for j := i + 1; j < len(points); j++ {
			end := points[j]
			path := 0
			for k := 0; k < len(points); k++ {
				p := points[k]
				if (end[0]-start[0])*(p[1]-start[1]) == (p[0]-start[0])*(end[1]-start[1]) {
					path++
				}
			}
			if path > max_path {
				max_path = path
			}
		}

	}

	return max_path
}

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}
