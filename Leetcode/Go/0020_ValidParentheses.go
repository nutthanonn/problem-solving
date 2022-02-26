package main

import "fmt"

func inArr(arr [3]string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func findIndex(arr [3]string, target string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func isValid(s string) bool {
	openString := [3]string{"(", "{", "["}
	closeString := [3]string{")", "}", "]"}
	q := []string{}
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	for _, v := range s {
		if inArr(openString, string(v)) {
			q = append(q, string(v))
		}

		if len(q) == 0 && inArr(closeString, string(v)) {
			return false
		}

		if len(q) > 0 && inArr(closeString, string(v)) && findIndex(openString, q[len(q)-1]) != findIndex(closeString, string(v)) {
			return false
		}

		if len(q) > 0 && inArr(closeString, string(v)) && findIndex(openString, q[len(q)-1]) == findIndex(closeString, string(v)) {
			q = q[:len(q)-1]
		}

	}

	if len(q) == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	s := "([}}])"
	fmt.Println(isValid(s))
}
