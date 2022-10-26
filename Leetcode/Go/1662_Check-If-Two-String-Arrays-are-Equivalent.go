package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
