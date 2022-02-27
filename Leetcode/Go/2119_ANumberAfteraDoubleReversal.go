package main

import (
	"fmt"
)

func isSameAfterReversals(num int) bool {
	isZero := num == 0
	isZeroBack := num%10 != 0
	return isZero || isZeroBack
}

func main() {
	number := 12500
	fmt.Println(isSameAfterReversals(number))
}
