package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pre := strs[0]
	for i := 0; i < len(pre); i++ {
		temp := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != temp {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	a := longestCommonPrefix([]string{"dog", "do", "dsa"})
	fmt.Println(a)
}
