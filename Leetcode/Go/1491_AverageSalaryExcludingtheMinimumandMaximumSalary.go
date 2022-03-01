package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {
	sort.Ints(salary)
	value := 0
	salary = salary[1 : len(salary)-1]
	for _, v := range salary {
		value += v
	}
	return float64(value) / float64(len(salary))
}

func main() {
	salary := []int{8000, 9000, 2000, 3000, 6000, 1000}
	fmt.Println(average(salary))
}
