package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	final := make(map[string][]string)
	for _, v := range strs {
		sorted := SortString(v)
		final[sorted] = append(final[sorted], v)
	}

	result := make([][]string, 0)
	for _, v := range final {
		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{""}))
}
