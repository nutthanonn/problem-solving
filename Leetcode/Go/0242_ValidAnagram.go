package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	x := []rune(s)
	x1 := []rune(t)
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	sort.Slice(x1, func(i, j int) bool { return x1[i] < x1[j] })
	if string(x) == string(x1) {
		return true
	}

	return false
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

}
