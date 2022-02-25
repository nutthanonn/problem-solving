package main

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	r := ""
	re := regexp.MustCompile(p)
	r2 := re.FindAllString(s, 1)
	for _, v := range r2 {
		r += v
	}
	return len(s) == len(r)

}

func main() {
	fmt.Println(isMatch("aa", "a"))
}
