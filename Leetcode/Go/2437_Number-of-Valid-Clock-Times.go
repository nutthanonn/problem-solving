package main

import (
	"fmt"
	"strings"
)

func countTime(time string) int {
	count := 1
	if time == "??:??" {
		return 24 * 60
	}

	time_arr := strings.Split(time, ":")

	hour := time_arr[0]
	minute := time_arr[1]

	if hour[0] == '?' {
		if hour[1] > '3' {
			count *= 2
		} else {
			count *= 3
		}
	}

	if hour[1] == '?' {
		if hour[0] == '2' {
			count *= 4
		} else {
			count *= 10
		}
	}

	if hour == "??" {
		count = 24
	}

	if minute[0] == '?' {
		count *= 6
	}

	if minute[1] == '?' {
		count *= 10
	}

	return count
}

func main() {
	fmt.Println(countTime("1?:??"))
}
