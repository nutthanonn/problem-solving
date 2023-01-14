package main

import "fmt"

func dp(start string, graph map[string][]rune, curr int, walk map[string]bool) rune {
	for _, v := range graph[start] {
		if curr > int(rune(v)) {
			curr = int(v)
		}

		if _, ok := walk[string(v)]; !ok {
			walk[string(v)] = true
			curr = int((dp(string(v), graph, curr, walk)))
		}
	}

	return rune(curr)
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	field := make(map[string][]rune)

	for i := 0; i < len(s1); i++ {
		field[string(s1[i])] = append(field[string(s1[i])], rune(s2[i]))
		field[string(s2[i])] = append(field[string(s2[i])], rune(s1[i]))
	}

	temp := ""
	for _, v := range baseStr {
		walk := make(map[string]bool)
		walk[string(v)] = true
		temp += string(dp(string(v), field, int(rune(v)), walk))
	}

	return temp
}

func main() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser"))
	fmt.Println(smallestEquivalentString("hello", "world", "hold"))
	fmt.Println(smallestEquivalentString("leetcode", "programs", "sourcecode"))
}
