package main

func addToArrayForm(num []int, k int) []int {
	var res []int
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		k /= 10
		if sum >= 10 {
			k++
			sum -= 10
		}
		res = append(res, sum)
	}

	for k > 0 {
		res = append(res, k%10)
		k /= 10
	}

	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}

func main() {
	addToArrayForm([]int{1, 2, 0, 0}, 34)
}
