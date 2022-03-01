package main

import "fmt"

func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}

func main() {
	low := 8
	high := 10
	fmt.Println(countOdds(low, high))
}
