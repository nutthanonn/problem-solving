package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	s := "aaaaa"
	fmt.Println(strStr(s, "bba"))
}
