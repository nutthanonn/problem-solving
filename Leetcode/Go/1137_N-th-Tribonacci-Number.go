package main

import "fmt"

func tribonacci(n int) int {
	field := make(map[int]int)
	field[0] = 0
	field[1] = 1
	field[2] = 1

	for i := 3; i < n+1; i++ {
		field[i] = field[i-3] + field[i-2] + field[i-1]
	}

	return field[n]
}

func main() {
	fmt.Println(tribonacci(4))
}
