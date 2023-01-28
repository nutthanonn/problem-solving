package main

import (
	"fmt"
	"strconv"
)

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minFlips(a int, b int, c int) int {
	binaryA := IntegerToBinary(a)
	binaryB := IntegerToBinary(b)
	binaryC := IntegerToBinary(c)
	maxLength := max(len(binaryA), max(len(binaryB), len(binaryC)))
	for len(binaryA) < maxLength {
		binaryA = "0" + binaryA
	}
	for len(binaryB) < maxLength {
		binaryB = "0" + binaryB
	}
	for len(binaryC) < maxLength {
		binaryC = "0" + binaryC
	}
	count := 0
	for i := 0; i < maxLength; i++ {
		if binaryC[i] == '0' {
			if binaryA[i] == '1' {
				count++
			}
			if binaryB[i] == '1' {
				count++
			}
		} else {
			if binaryA[i] == '0' && binaryB[i] == '0' {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(minFlips(2, 6, 5))
	fmt.Println(minFlips(4, 2, 7))
	fmt.Println(minFlips(1, 2, 3))
	fmt.Println(minFlips(8, 3, 5))
}
