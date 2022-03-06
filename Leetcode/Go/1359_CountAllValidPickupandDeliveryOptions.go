package main

import "fmt"

func countOrders(n int) int {
	ans, m := 1, 1000000007
	for i := 2; i <= n; i++ {
		ans *= (2*i - 1) * i
		ans %= m
	}
	return ans
}

func main() {
	number := 3
	fmt.Println(countOrders(number))
}
