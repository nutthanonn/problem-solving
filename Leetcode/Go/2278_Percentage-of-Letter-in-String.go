package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	count := 0
	for _, i := range s {
		if i == rune(letter) {
			count++
		}
	}
	return count * 100 / len(s)
}

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
}
