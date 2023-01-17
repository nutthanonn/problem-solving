package main

import "fmt"

func comb(start string, field map[int][]rune, digits string) []string {
	if len(digits) == 0 {
		return []string{start}
	}

	var result []string

	for _, r := range field[int(digits[0]-'0')] {
		result = append(result, comb(start+string(r), field, digits[1:])...)
	}

	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	fields := make(map[int][]rune)

	fields[2] = []rune{'a', 'b', 'c'}
	fields[3] = []rune{'d', 'e', 'f'}
	fields[4] = []rune{'g', 'h', 'i'}
	fields[5] = []rune{'j', 'k', 'l'}
	fields[6] = []rune{'m', 'n', 'o'}
	fields[7] = []rune{'p', 'q', 'r', 's'}
	fields[8] = []rune{'t', 'u', 'v'}
	fields[9] = []rune{'w', 'x', 'y', 'z'}

	return comb("", fields, digits)
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}
