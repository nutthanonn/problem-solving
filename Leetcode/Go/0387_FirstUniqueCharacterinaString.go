package main

import "fmt"

func firstUniqChar(s string) int {
	x := make(map[string]int)
	for _, v := range s {
		x[string(v)]++
	}
	for n, v := range s {
		if x[string(v)] == 1 {
			return n
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
