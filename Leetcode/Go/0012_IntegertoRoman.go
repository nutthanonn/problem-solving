package main

import "fmt"

func intToRoman(num int) string {
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for num > 0 {
		for i := range value {
			if num >= value[i] {
				num -= value[i]
				result += symbol[i]
				break
			}
		}
	}
	return result

}

func main() {
	number := 1000
	fmt.Println(intToRoman(number))
}
