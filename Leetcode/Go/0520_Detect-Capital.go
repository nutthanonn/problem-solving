package main

import "fmt"

func allLower(word string) bool {
	for _, v := range word {
		if int(v) >= 65 && int(v) <= 90 {
			return false
		}
	}
	return true
}

func allUpper(word string) bool {
	for _, v := range word {
		if !(int(v) >= 65 && int(v) <= 90) {
			return false
		}
	}
	return true
}

func isCapi(word string) bool {
	if int(rune(word[0])) >= 65 && int(rune(word[0])) <= 90 {
		for i := 1; i < len(word); i++ {
			return allLower(word[1:])
		}
	}

	return false
}

func detectCapitalUse(word string) bool {
	return isCapi(word) || allUpper(word) || allLower(word)
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("leetcode"))
	fmt.Println(detectCapitalUse("Google"))
	fmt.Println(detectCapitalUse("FlaG"))
	fmt.Println(detectCapitalUse("ffffffffffffffffffffF"))
}
