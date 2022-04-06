package main

import (
	"fmt"
	"math"
	"strings"
)

func reverseWords(s string) string {
	b := []byte(s)
	n := len(s)
	for i := 0; i < int(math.Floor(float64(n/2))); i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	arr := strings.Split(string(b), " ")

	n = len(arr)
	for i := 0; i < int(math.Floor(float64(n/2))); i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	return strings.Join(arr, " ")
}

func main() {
	fmt.Println(reverseWords("God Ding"))
}
