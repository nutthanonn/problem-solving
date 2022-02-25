package main

import (
	"fmt"
	"strconv"
)

func myAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	return i
}

func main() {
	fmt.Println(myAtoi("-00042"))
}
