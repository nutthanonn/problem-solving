package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1, arr2 := make([]int, 26), make([]int, 26)
	s1, s2 := [26]int{}, [26]int{}
	for i := range word1 {
		arr1[word1[i]-'a']++
		arr2[word2[i]-'a']++
		s1[word1[i]-'a'] = 1
		s2[word2[i]-'a'] = 1
	}

	sort.Ints(arr1)
	sort.Ints(arr2)
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return s1 == s2
}

func main() {
	fmt.Println(closeStrings("abc", "abc"))
}
