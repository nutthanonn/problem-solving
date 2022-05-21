package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	l := len(s)
	c := false
	if strings.Index(s, "a") == -1 {
		return true
	}
	for i := 0; i < l; i++ {
		if s[i] == 'b' {
			c = true
		}

		if c && s[i] == 'a' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkString("bbb"))
}
