package main

func arraySign(nums []int) int {
	final := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}

		if v < 0 {
			final *= -1
		}
	}

	return final
}

func main() {

}
