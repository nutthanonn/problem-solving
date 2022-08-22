package main

import "fmt"

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	for {
		if n == 1 || n == 4 {
			return true
		}

		if n%4 != 0 {
			return false
		}
		n /= 4
	}
}

func main() {
	fmt.Println(isPowerOfFour(4))
}
