package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	var (
		maxLen int
		start  int
	)

	seen := make(map[rune]int)
	for i, r := range s {
		if j, ok := seen[r]; ok {
			start = max(start, j+1)
		}
		maxLen = max(maxLen, i-start+1)
		seen[r] = i
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("ddddd"))
}
