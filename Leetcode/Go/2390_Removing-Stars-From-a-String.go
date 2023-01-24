package main

import "fmt"

func removeStars(s string) string {
	l := len(s)
	txt := ""
	for i := 0; i < l; i++ {
		if s[i] == '*' {
			txt = txt[:len(txt)-1]
		} else {
			txt += string(s[i])
		}
	}
	return txt
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))

}
