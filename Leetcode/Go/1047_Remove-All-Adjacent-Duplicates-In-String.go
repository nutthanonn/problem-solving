package main

import (
	"fmt"
)

func removeDuplicates(s string) string {
	final := ""

	for i := 0; i < len(s); i++ {
		if len(final) > 0 && final[len(final)-1] == s[i] {
			final = final[:len(final)-1]
		} else {
			final += string(s[i])
		}
	}

	return final
}

func main() {
	fmt.Println(removeDuplicates("azxxzy"))
}
