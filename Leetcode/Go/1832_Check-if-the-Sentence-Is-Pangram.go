package main

import "fmt"

func checkIfPangram(sentence string) bool {
	bucket := [26]int{0}
	for _, v := range sentence {
		bucket[v-97] += 1
	}

	for _, v := range bucket {
		if v == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))

}
