package main

import "fmt"

func isValid(s string) bool {
	field := make(map[rune]rune)
	field[')'] = '('
	field[']'] = '['
	field['}'] = '{'

	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if field[v] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("([}}])"))
	fmt.Println(isValid("()[]{}"))
}
