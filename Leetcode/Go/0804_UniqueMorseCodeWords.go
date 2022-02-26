package main

import "fmt"

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func uniqueMorseRepresentations(words []string) int {
	alphabet := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	diff := []string{}
	for _, v := range words {
		txt := ""
		for _, j := range v {
			txt += alphabet[j-97]
		}
		diff = append(diff, txt)
	}
	return len(diff)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
